package engine

import (
	"context"
	"github.com/autom8ter/engine/config"
	"github.com/autom8ter/engine/servers"
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/grpclog"
	"os"
	"os/signal"
	"syscall"
)

// Engine is an interface used to describe a server runtime
type Engine interface {
	With(opts ...config.Option) *Runtime
	Shutdown()
	Serve(ctx context.Context) error
}

// New creates a new engine intstance.
func New(network, addr string) Engine {
	return &Runtime{
		cfg: config.New(network, addr),
	}
}

// Runtime is an implementation of the engine API.
type Runtime struct {
	ctx        context.Context
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

// Serve starts the runtime gRPC server.
func (e *Runtime) Serve(ctx context.Context) error {
	var err error
	eg, ctx := errgroup.WithContext(ctx)
	ctx, e.cancelFunc = context.WithCancel(ctx)

	grpcServer := servers.NewGrpcServer(e.cfg)
	groclis := e.cfg.GRPC()

	httpServer := servers.NewHTTPServer(e.cfg)
	httpLis := e.cfg.HTTP()

	{
		grpclog.Infoln("starting grpc server....")
		eg.Go(func() error {
			return grpcServer.Serve(groclis)
		})
		grpclog.Infoln("starting http server....")
		eg.Go(func() error {
			return httpServer.Serve(httpLis)
		})
		grpclog.Infoln("starting multiplex server....")
		eg.Go(e.cfg.Serve)
		grpclog.Infoln("press ctr-c to shutdown")
		eg.Go(func() error {
			return e.watchShutdownSignal(ctx)
		})
	}
	err = eg.Wait()
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
