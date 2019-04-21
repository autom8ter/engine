package engine

import (
	"context"
	"fmt"
	"github.com/autom8ter/engine/config"
	"github.com/autom8ter/engine/driver"
	"github.com/autom8ter/engine/servers"
	"github.com/autom8ter/objectify"
	"github.com/pkg/errors"
	"google.golang.org/grpc/grpclog"
)

var tool = objectify.New()

// Engine is an interface used to describe a server runtime
type Engine interface {
	With(opts ...config.Option) *Runtime
	Config() *config.Config
	Shutdown(ctx context.Context)
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

// Runtime is an implementation of the engine API.
type Runtime struct {
	cfg        *config.Config `validate:"required"`
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
	if err := tool.Validate(e); err != nil {
		return tool.WrapErrf(err, "Method: %s", "engine.Serve")
	}
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
func (e *Runtime) Shutdown(ctx context.Context) {
	if ctx == nil {
		ctx = context.Background()
	}
	_ = tool.WatchForShutdown(ctx, e.cancelFunc)
}

//Serve creates starts a gRPC server without the need to create an engine instance
func Serve(addr string, debug bool, plugs ...driver.Plugin) error {
	return New("tcp", addr, debug).With(config.WithDefaultPlugins(), config.WithDefaultMiddlewares(), config.WithPlugins(plugs...)).Serve()
}
