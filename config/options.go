package config

import (
	"github.com/autom8ter/engine/driver"
	"github.com/autom8ter/engine/lib/util"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"strings"
)

// Option configures a gRPC and a gateway server.
type Option func(*Config)

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

// WithGoPlugins returns an Option that adds hard-coded Plugins(golang) to the engine runtime as opposed to go/plugins.
func WithGoPlugins(svrs ...driver.Plugin) Option {
	return func(c *Config) {
		c.Plugins = append(c.Plugins, svrs...)
		if len(c.Plugins) == 0 {
			grpclog.Fatal(errors.New("zero valid plugins registered"))
		}
	}
}

// WithGoPlugins returns an Option that adds hard-coded Plugins(golang) to the engine runtime as opposed to go/plugins.
func WithPluginSymbol(sym string) Option {
	return func(c *Config) {
		c.Symbol = sym
	}
}

// WithGoPlugins returns an Option that adds hard-coded Plugins(golang) to the engine runtime as opposed to go/plugins.
func WithEnvPrefix(prefix string) Option {
	return func(c *Config) {
		util.Debugf("loading environmental variables with prefix: %s\n", prefix)
		viper.SetEnvPrefix(prefix)
		util.Debugf("setting environmental key replacer: replace: %s with: %s\n", ".", "_")
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
		viper.AutomaticEnv()
	}
}
