package config

import (
	"github.com/autom8ter/engine/driver"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"os"
)

func init() {
	grpclog.SetLoggerV2(grpclog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout))
}

// Option configures a gRPC and a gateway server.
type Option func(*Config)

//With is used to configure/initialize a Config with custom options
func (c *Config) With(opts []Option) *Config {
	for _, f := range opts {
		f(c)
	}
	return c
}

// WithGoPlugins returns an Option that adds hard-coded Plugins(golang) to the engine runtime as opposed to go/plugins.
func WithGoPlugins(svrs ...driver.Plugin) Option {
	return func(c *Config) {
		c.Plugins = append(c.Plugins, svrs...)
		if len(c.Plugins) == 0 {
			grpclog.Fatal(errors.New("zero valid plugins registered"))
		}
	}
}

// WithNetwork returns an Option that sets an network address for a gRPC server.
func WithNetwork(network, addr string) Option {
	viper.Set("network", network)
	viper.Set("address", addr)
	return func(c *Config) {
		c.Address = addr
		c.Network = network
	}
}

// WithUnaryInterceptors returns an Option that sets unary interceptor(s) for a gRPC server.
func WithUnaryInterceptors(interceptors ...grpc.UnaryServerInterceptor) Option {
	return func(c *Config) {
		c.UnaryInterceptors = append(c.UnaryInterceptors, interceptors...)
	}
}

// WithGrpcServerStreamInterceptors returns an Option that sets stream interceptor(s) for a gRPC server.
func WithStreamInterceptors(interceptors ...grpc.StreamServerInterceptor) Option {
	return func(c *Config) {
		c.StreamInterceptors = append(c.StreamInterceptors, interceptors...)
	}
}

// WithOptions returns an Option that sets grpc.ServerOption(s) to a gRPC server.
func WithServerOptions(opts ...grpc.ServerOption) Option {
	return func(c *Config) {
		c.Option = append(c.Option, opts...)
	}
}

// WithPluginPaths adds relative filepaths to Plugins to add to the engine runtime
func WithPluginPaths(paths ...string) Option {
	return func(c *Config) {
		c.Paths = append(c.Paths, paths...)
	}
}
