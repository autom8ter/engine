package engine

import (
	"context"
	"github.com/autom8ter/engine/config"
	"github.com/autom8ter/engine/driver"
	"github.com/pkg/errors"
	"github.com/soheilhy/cmux"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc/grpclog"
	"net"
	"os"
	"os/signal"
	"reflect"
	"syscall"
)

// Engine is the framework instance.
type Engine struct {
	*config.Config
	cancelFunc func()
}

// New creates a server intstance.
func (e *Engine) With(opts ...config.Option) *Engine {
	return &Engine{
		Config: e.Config.With(opts),
	}
}

// New creates a server intstance.
func New(plugins ...driver.Plugin) *Engine {
	return &Engine{
		Config: config.New(plugins...),
	}
}



// Serve starts gRPC and Gateway servers.
func (e *Engine) Serve() error {
	var (
		grpcServer, gatewayServer, muxServer driver.Engine
		grpcLis, gatewayLis, internalLis     net.Listener
		err                                  error
	)

	if e.GrpcAddr != nil && e.GatewayAddr != nil && reflect.DeepEqual(e.GrpcAddr, e.GatewayAddr) {
		lis, err := e.GrpcAddr.CreateListener()
		if err != nil {
			return errors.Wrap(err, "failed to listen network for servers")
		}
		mux := cmux.New(lis)
		muxServer = NewMuxServer(mux, lis)
		grpcLis = mux.Match(
			cmux.HTTP2HeaderField("content-type", "Runtime/grpc"),
			cmux.HTTP2HeaderField("content-type", "application/grpc"),
			cmux.HTTP2HeaderField("content-type", "grpc"),
		)
		gatewayLis = mux.Match(cmux.HTTP2(), cmux.HTTP1Fast())
	}

	// Setup servers
	grpcServer = NewGrpcServer(e.Config)

	// Setup listeners
	if grpcLis == nil && e.GrpcAddr != nil {
		grpcLis, err = e.GrpcAddr.CreateListener()
		if err != nil {
			return errors.Wrap(err, "failed to listen network for gRPC server")
		}
		defer grpcLis.Close()
	}

	if e.GatewayAddr != nil {
		gatewayServer = NewGatewayServer(e.Config)
		internalLis, err = e.GrpcInternalAddr.CreateListener()
		if err != nil {
			return errors.Wrap(err, "failed to listen network for gRPC server internal")
		}
		defer internalLis.Close()
	}

	if gatewayLis == nil && e.GatewayAddr != nil {
		gatewayLis, err = e.GatewayAddr.CreateListener()
		if err != nil {
			return errors.Wrap(err, "failed to listen network for gateway server")
		}
		defer gatewayLis.Close()
	}

	// Start servers
	eg, ctx := errgroup.WithContext(context.Background())
	ctx, e.cancelFunc = context.WithCancel(ctx)

	if internalLis != nil {
		eg.Go(func() error { return grpcServer.Serve(internalLis) })
	}
	if grpcLis != nil {
		eg.Go(func() error { return grpcServer.Serve(grpcLis) })
	}
	if gatewayLis != nil {
		eg.Go(func() error { return gatewayServer.Serve(gatewayLis) })
	}
	if muxServer != nil {
		eg.Go(func() error { return muxServer.Serve(nil) })
	}

	eg.Go(func() error { return e.watchShutdownSignal(ctx) })

	select {
	case <-ctx.Done():
		for _, s := range []driver.Engine{gatewayServer, grpcServer, muxServer} {
			if s != nil {
				s.Shutdown()
			}
		}
	}

	err = eg.Wait()

	return errors.WithStack(err)
}

// Shutdown closes servers.
func (e *Engine) Shutdown() {
	if e.cancelFunc != nil {
		e.cancelFunc()
	} else {
		grpclog.Warning("the server has been started yet")
	}
}

func (e *Engine) watchShutdownSignal(ctx context.Context) error {
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
