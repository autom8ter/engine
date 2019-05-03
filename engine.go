package engine

import (
	"context"
	"github.com/autom8ter/engine/config"
	"github.com/autom8ter/engine/driver"
	"github.com/autom8ter/engine/lis"
	"github.com/autom8ter/engine/servers"
	"github.com/autom8ter/objectify"
	"github.com/pkg/errors"
	"google.golang.org/grpc/grpclog"
	"net/http"
)

var tool = objectify.Default()

// Engine is an interface used to describe a server runtime
type Engine interface {
	With(opts ...config.Option) *Runtime
	Config() *config.Config
	Shutdown(ctx context.Context)
	Serve() error
}

// New creates a new engine intstance.
func New(network, addr string, debug bool) Engine {
	cfg := config.New(network, addr, debug)
	l, err := cfg.CreateListener()
	if err != nil {
		grpclog.Fatalln(err.Error())
	}
	return &Runtime{
		cfg: cfg,
		lis: l,
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
	lis        *lis.Listener
	cfg        *config.Config `validate:"required"`
	cancelFunc []func()
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
	e.cancelFunc = append(e.cancelFunc, grpcServer.Shutdown)
	return errors.WithStack(grpcServer.Serve(e.lis.GRPCListener()))
}

// Serve starts the runtime gRPC server.
func (e *Runtime) Proxy(server *http.Server) error {
	e.cancelFunc = append(e.cancelFunc, func() {
		if err := server.Shutdown(context.TODO()); err != nil {
			grpclog.Infoln(err.Error())
		}
	})

	return errors.WithStack(server.Serve(e.lis.HTTPListener()))
}

// Shutdown gracefully closes the grpc server.
func (e *Runtime) Shutdown(ctx context.Context) {
	if ctx == nil {
		ctx = context.Background()
	}
	_ = tool.WatchForShutdown(ctx, func() {
		for _, c := range e.cancelFunc {
			c()
		}
	})
}

//Serve creates starts a gRPC server without the need to create an engine instance
func Serve(addr string, debug bool, plugs ...driver.Plugin) error {
	return New("tcp", addr, debug).With(config.WithDefaultPlugins(), config.WithDefaultMiddlewares(), config.WithPlugins(plugs...)).Serve()
}
