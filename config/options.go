package config

import (
	"fmt"
	"github.com/autom8ter/engine/driver"
	"github.com/autom8ter/engine/handlers"
	"github.com/autom8ter/engine/plugin"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"log"
	"os"
)

// Option configures a gRPC and a gateway server.
type Option func(*Config)

func (c *Config) With(opts []Option) *Config {
	for _, f := range opts {
		f(c)
	}
	return c
}

// WithServers returns an Option that sets gRPC service server implementation(s).
func WithGoPlugins(svrs ...driver.Plugin) Option {
	return func(c *Config) {
		c.Plugins = append(c.Plugins, svrs...)
		if len(c.Plugins) == 0 {
			log.Fatal(errors.New("zero valid plugins registered"))
		}
	}
}

// WithAddr returns an Option that sets an network address for a gRPC and a gateway server.
func WithAddr(network, addr string) Option {
	return func(c *Config) {
		WithGrpcAddr(network, addr)(c)
		WithGatewayAddr(network, addr)(c)
	}
}

// WithGrpcAddr returns an Option that sets an network address for a gRPC server.
func WithGrpcAddr(network, addr string) Option {
	address := &Address{
		Network: network,
		Addr:    addr,
	}
	viper.Set("grpc", address)
	return func(c *Config) {
		c.GrpcAddr = &Address{
			Network: network,
			Addr:    addr,
		}
	}
}

// WithGrpcInternalAddr returns an Option that sets an network address connected by a gateway server.
func WithGrpcInternalAddr(network, addr string) Option {
	address := &Address{
		Network: network,
		Addr:    addr,
	}
	viper.Set("internal", address)
	return func(c *Config) {
		c.GrpcInternalAddr = address
	}
}

// WithGatewayAddr returns an Option that sets an network address for a gateway server.
func WithGatewayAddr(network, addr string) Option {
	address := &Address{
		Network: network,
		Addr:    addr,
	}
	viper.Set("gateway", address)
	return func(c *Config) {
		c.GatewayAddr = &Address{
			Network: network,
			Addr:    addr,
		}
	}
}

// WithGrpcServerUnaryInterceptors returns an Option that sets unary interceptor(s) for a gRPC server.
func WithGrpcServerUnaryInterceptors(interceptors ...grpc.UnaryServerInterceptor) Option {
	return func(c *Config) {
		c.GrpcServerUnaryInterceptors = append(c.GrpcServerUnaryInterceptors, interceptors...)
	}
}

// WithGrpcServerStreamInterceptors returns an Option that sets stream interceptor(s) for a gRPC server.
func WithGrpcServerStreamInterceptors(interceptors ...grpc.StreamServerInterceptor) Option {
	return func(c *Config) {
		c.GrpcServerStreamInterceptors = append(c.GrpcServerStreamInterceptors, interceptors...)
	}
}

// WithGatewayServerUnaryInterceptors returns an Option that sets unary interceptor(s) for a gRPC client used by a gateway server.
func WithGatewayServerUnaryInterceptors(interceptors ...grpc.UnaryClientInterceptor) Option {
	return func(c *Config) {
		c.GatewayServerUnaryInterceptors = append(c.GatewayServerUnaryInterceptors, interceptors...)
	}
}

// WithGatewayServerStreamInterceptors returns an Option that sets stream interceptor(s) for a gRPC client used by a gateway server.
func WithGatewayServerStreamInterceptors(interceptors ...grpc.StreamClientInterceptor) Option {
	return func(c *Config) {
		c.GatewayServerStreamInterceptors = append(c.GatewayServerStreamInterceptors, interceptors...)
	}
}

// WithGrpcServerOptions returns an Option that sets grpc.ServerOption(s) to a gRPC server.
func WithGrpcServerOptions(opts ...grpc.ServerOption) Option {
	return func(c *Config) {
		c.GrpcServerOption = append(c.GrpcServerOption, opts...)
	}
}

// WithGatewayDialOptions returns an Option that sets grpc.DialOption(s) to a gRPC clinet used by a gateway server.
func WithGatewayDialOptions(opts ...grpc.DialOption) Option {
	return func(c *Config) {
		c.GatewayDialOption = append(c.GatewayDialOption, opts...)
	}
}

// WithGatewayMuxOptions returns an Option that sets runtime.ServeMuxOption(s) to a gateway server.
func WithGatewayMuxOptions(opts ...runtime.ServeMuxOption) Option {
	return func(c *Config) {
		c.GatewayMuxOptions = append(c.GatewayMuxOptions, opts...)
	}
}

// WithGatewayServerMiddlewares returns an Option that sets middleware(s) for http.driver.Plugin to a gateway server.
func WithGatewayServerMiddlewares(middlewares ...handlers.HTTPServerMiddleware) Option {
	return func(c *Config) {
		c.GatewayServerMiddlewares = append(c.GatewayServerMiddlewares, middlewares...)
	}
}

// WithGatewayServerConfig returns an Option that specifies http.driver.Plugin configuration to a gateway server.
func WithGatewayServerConfig(cfg *HTTPServerConfig) Option {
	return func(c *Config) {
		c.GatewayServerConfig = cfg
	}
}

// WithPassedHeader returns an Option that sets configurations about passed headers for a gateway server.
func WithPassedHeader(decider handlers.PassedHeaderDeciderFunc) Option {
	return WithGatewayServerMiddlewares(handlers.CreatePassingHeaderMiddleware(decider))
}

// WithDefaultLogger returns an Option that sets default grpclogger.LoggerV2 object.
func WithDefaultLogger() Option {
	return func(c *Config) {
		grpclog.SetLoggerV2(grpclog.NewLoggerV2(os.Stdout, os.Stderr, os.Stderr))
	}
}

// WithPassedHeader returns an Option that sets configurations about passed headers for a gateway server.
func WithSwaggerFile(path string) Option {
	viper.Set("swagger", path)
	return func(config *Config) {
		config.Swagger = path
	}
}

func WithPlugins() Option {
	return func(config *Config) {
		config.Plugins = append(config.Plugins, plugin.LoadPlugins()...)
		if len(config.Plugins) == 0 {
			errStr := fmt.Sprintf("zero valid plugins registered. plugin files: %s", plugin.Files())
			log.Fatal(errors.New(errStr))
		}
	}
}
