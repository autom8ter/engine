package engine

import (
	"context"
	"github.com/autom8ter/engine/config"
	"github.com/autom8ter/engine/servers"
	"github.com/pkg/errors"
	"google.golang.org/grpc/grpclog"
	"os"
	"os/signal"
	"syscall"
)

type Engine interface {
	With(opts ...config.Option) *Runtime
	Config() *config.Config
	Serve() error
	Shutdown()
}

// Engine is the framework instance.
type Runtime struct {
	cfg        *config.Config
	cancelFunc func()
}

// New creates a engine intstance.
func New() Engine {
	return &Runtime{
		cfg: config.New(),
	}
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
	lis, err := e.cfg.CreateListener()
	if err != nil {
		grpclog.Fatal(err.Error())
	}
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
