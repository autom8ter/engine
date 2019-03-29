package engine

import (
	"context"
	"fmt"
	"github.com/autom8ter/engine/config"
	"github.com/autom8ter/engine/driver"
	"github.com/autom8ter/engine/servers"
	"github.com/pkg/errors"
	"google.golang.org/grpc/grpclog"
	"os"
	"os/signal"
	"syscall"
)

// Engine is an interface used to describe a server runtime
type Engine interface {
	With(opts ...config.Option) *Runtime
	Config() *config.Config
	Shutdown()
	Serve() error
}

// New creates a new engine intstance.
func New(network, addr string, debug bool) Engine {
	return &Runtime{
		cfg: config.New(network, addr, debug),
	}
}

// New creates a new engine intstance.
func Default(network, addr string, debug bool) Engine {
	r := &Runtime{
		cfg: config.New(network, addr, debug),
	}
	r.With(config.WithDefaultPlugins(), config.WithDefaultMiddlewares())
	return r
}

// New creates a new engine intstance.
func GetRuntime(network, addr string, debug bool) *Runtime {
	return &Runtime{
		cfg:        config.New(network, addr, debug).With(config.WithDefaultPlugins(), config.WithDefaultMiddlewares()),
		cancelFunc: nil,
	}
}

// Runtime is an implementation of the engine API.
type Runtime struct {
	cfg        *config.Config
	cancelFunc func()
}

// With wraps the runtimes config with config options
// ref: github.com/autom8ter/engine/config/options.go
func (e *Runtime) With(opts ...config.Option) *Runtime {
	return &Runtime{
		cfg: e.cfg.With(opts...),
	}
}

// Config returns the runtimes current configuration
func (e *Runtime) Config() *config.Config {
	return e.cfg
}

// Serve starts the runtime gRPC server.
func (e *Runtime) Serve() error {
	grpcServer := servers.NewGrpcServer(e.cfg)
	e.cancelFunc = grpcServer.Shutdown
	lis, err := e.cfg.CreateListener()
	if err != nil {
		grpclog.Fatalln(err.Error())
	}
	fmt.Println(fmt.Sprintln(e.cfg.Debug(), len(e.cfg.UnaryInterceptors), len(e.cfg.StreamInterceptors), len(e.cfg.Option), len(e.cfg.Plugins), e.cfg.Network, e.cfg.Address))
	err = grpcServer.Serve(lis)
	return errors.WithStack(err)
}

// Shutdown gracefully closes the grpc server.
func (e *Runtime) Shutdown() {
	e.cancelFunc()
}

func (e *Runtime) watchShutdownSignal(ctx context.Context) error {
	sdCh := make(chan os.Signal, 1)
	defer close(sdCh)
	defer signal.Stop(sdCh)
	signal.Notify(sdCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-sdCh:
		e.Shutdown()
	case <-ctx.Done():
		// no-op
	}
	return nil
}

func Serve(addr string, debug bool, plugs ...driver.Plugin) error {
	return New("tcp", addr, debug).With(config.WithDefaultPlugins(), config.WithDefaultMiddlewares(), config.WithPlugins(plugs...)).Serve()
}
