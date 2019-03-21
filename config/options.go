package config

import (
	"github.com/autom8ter/engine/driver"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
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
			grpclog.Fatal(errors.New("zero valid plugins registered"))
		}
	}
}

// WithGrpcAddr returns an Option that sets an network address for a gRPC server.
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
func WithOptions(opts ...grpc.ServerOption) Option {
	return func(c *Config) {
		c.Option = append(c.Option, opts...)
	}
}

// WithDefaultLogger returns an Option that sets default grpclogger.LoggerV2 object.
func WithGRPCLogger() Option {
	return func(c *Config) {
		grpclog.SetLoggerV2(grpclog.NewLoggerV2(os.Stdout, os.Stderr, os.Stderr))
	}
}
