package engine

import (
	"context"
	"fmt"
	"github.com/autom8ter/engine/config"
	"github.com/autom8ter/engine/driver"
	"github.com/autom8ter/engine/listeners"
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
	cfg        *config.Config
	cancelFunc func()
}

// New creates a server intstance.
func (e *Engine) With(opts ...config.Option) *Engine {
	return &Engine{
		cfg: e.cfg.With(opts),
	}
}

// New creates a server intstance.
func (e *Engine) Config() *config.Config {
	return e.cfg
}

// New creates a server intstance.
func New(plugins ...driver.Plugin) *Engine {
	return &Engine{
		cfg: config.New(plugins...),
	}
}

// Serve starts gRPC and Gateway servers.
func (e *Engine) Serve() error {
	var (
		grpcServer, gatewayServer, muxProxy driver.Listener
		grpcLis, gatewayLis, internalLis    net.Listener
		err                                 error
	)

	if e.cfg.GrpcAddr != nil && e.cfg.GatewayAddr != nil && reflect.DeepEqual(e.cfg.GrpcAddr, e.cfg.GatewayAddr) {
		lis, err := e.cfg.GrpcAddr.CreateListener()
		if err != nil {
			return errors.Wrap(err, "failed to listen network for servers")
		}
		mux := cmux.New(lis)
		muxProxy = listeners.NewMuxProxy(mux, lis)
		grpcLis = mux.Match(
			cmux.HTTP2HeaderField("content-type", "Runtime/grpc"),
			cmux.HTTP2HeaderField("content-type", "application/grpc"),
			cmux.HTTP2HeaderField("content-type", "grpc"),
		)
		gatewayLis = mux.Match(cmux.HTTP2(), cmux.HTTP1Fast())
	}

	// Setup servers
	grpcServer = listeners.NewGrpcListener(e.cfg)

	// Setup listeners
	if grpcLis == nil && e.cfg.GrpcAddr != nil {
		grpcLis, err = e.cfg.GrpcAddr.CreateListener()
		if err != nil {
			return errors.Wrap(err, "failed to listen network for gRPC server")
		}
		defer grpcLis.Close()
	}

	if e.cfg.GatewayAddr != nil {
		gatewayServer = listeners.NewGatewayListener(e.cfg)
		internalLis, err = e.cfg.GrpcInternalAddr.CreateListener()
		if err != nil {
			return errors.Wrap(err, "failed to listen network for gRPC server internal")
		}
		defer internalLis.Close()
	}

	if gatewayLis == nil && e.cfg.GatewayAddr != nil {
		gatewayLis, err = e.cfg.GatewayAddr.CreateListener()
		if err != nil {
			return errors.Wrap(err, "failed to listen network for gateway server")
		}
		defer gatewayLis.Close()
	}

	// Start servers
	eg, ctx := errgroup.WithContext(context.Background())
	ctx, e.cancelFunc = context.WithCancel(ctx)
	fmt.Println(fmt.Sprintf("registered %v plugin(s)", len(e.cfg.Plugins)))
	if internalLis != nil {
		eg.Go(func() error { return grpcServer.Serve(internalLis) })
	}
	if grpcLis != nil {
		eg.Go(func() error { return grpcServer.Serve(grpcLis) })
	}
	if gatewayLis != nil {
		eg.Go(func() error { return gatewayServer.Serve(gatewayLis) })
	}
	if muxProxy != nil {
		eg.Go(func() error { return muxProxy.Serve(nil) })
	}

	eg.Go(func() error { return e.watchShutdownSignal(ctx) })

	select {
	case <-ctx.Done():
		for _, s := range []driver.Listener{gatewayServer, grpcServer, muxProxy} {
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
