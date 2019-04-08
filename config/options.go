package config

import (
	"github.com/autom8ter/engine/driver"
	"github.com/autom8ter/engine/lib"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/stats"
	"time"
)

// Option configures a gRPC and a gateway server.
type Option func(*Config)

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

// WithGoPlugins returns an Option that adds hard-coded Plugins(golang) to the engine runtime as opposed to go/plugins.
// See driver.Plugin for the interface definition.
func WithPlugins(svrs ...driver.Plugin) Option {
	return func(c *Config) {
		c.Plugins = append(c.Plugins, svrs...)
		if len(c.Plugins) == 0 {
		}
	}
}

// WithCreds returns a ServerOption that sets credentials for server connections.
func WithCreds(creds credentials.TransportCredentials) Option {
	return func(c *Config) {
		c.Option = append(c.Option, grpc.Creds(creds))
	}
}

// WithStatsHandler returns a ServerOption that sets the stats handler for the server.
func WithStatsHandler(h stats.Handler) Option {
	return func(c *Config) {
		c.Option = append(c.Option, grpc.StatsHandler(h))
	}
}

// WithStatsHandler ConnectionTimeout returns a ServerOption that sets the timeout for connection establishment (up to and including HTTP/2 handshaking) for all new connections.
// If this is not set, the default is 120 seconds.
func WithConnTimeout(t time.Duration) Option {
	return func(c *Config) {
		c.Option = append(c.Option, grpc.ConnectionTimeout(t))
	}
}

// WithMaxConcurrentStreams returns a ServerOption that will apply a limit on the number
// of concurrent streams to each ServerTransport.
func WithMaxConcurrentStreams(num uint32) Option {
	return func(c *Config) {
		c.Option = append(c.Option, grpc.MaxConcurrentStreams(num))
	}
}

// WithChannelz adds grpc server channelz to the list of plugins ref: https://godoc.org/google.golang.org/grpc/channelz/grpc_channelz_v1
func WithChannelz() Option {
	return func(c *Config) {
		c.Plugins = append(c.Plugins, lib.NewChannelz())
	}
}

// WithReflection adds grpc server reflection to the list of plugins ref: https://godoc.org/google.golang.org/grpc/reflection
func WithReflection() Option {
	return func(c *Config) {
		c.Plugins = append(c.Plugins, lib.NewReflection())
	}
}

// WithHealthz exposes server's health and it must be imported to enable support for client-side health checks and adds it to plugins. ref: https://godoc.org/google.golang.org/grpc/health
func WithHealthz() Option {
	return func(c *Config) {
		c.Plugins = append(c.Plugins, lib.NewHealthz())
	}
}

func WithUnaryLoggingMiddleware() Option {
	return func(c *Config) {
		c.UnaryInterceptors = append(c.UnaryInterceptors, lib.NewUnaryLogger())
	}
}

func WithStreamLoggingMiddleware() Option {

	return func(c *Config) {
		c.StreamInterceptors = append(c.StreamInterceptors, lib.NewStreamLogger())
	}
}

func WithUnaryRecoveryMiddleware() Option {
	return func(c *Config) {
		c.UnaryInterceptors = append(c.UnaryInterceptors, grpc_recovery.UnaryServerInterceptor())
	}
}

func WithStreamRecoveryMiddleware() Option {
	return func(c *Config) {
		c.StreamInterceptors = append(c.StreamInterceptors, grpc_recovery.StreamServerInterceptor())
	}
}

func WithUnaryTraceMiddleware() Option {

	return func(c *Config) {
		c.UnaryInterceptors = append(c.UnaryInterceptors, grpc_opentracing.UnaryServerInterceptor())
	}
}

func WithStreamTraceMiddleware() Option {
	return func(c *Config) {
		c.StreamInterceptors = append(c.StreamInterceptors, grpc_opentracing.StreamServerInterceptor())
	}
}

func WithUnaryUUIDMiddleware() Option {
	return func(c *Config) {
		c.UnaryInterceptors = append(c.UnaryInterceptors, lib.NewUnaryUUID())
	}
}

func WithDefaultMiddlewares() Option {
	return func(c *Config) {
		c.UnaryInterceptors = append(c.UnaryInterceptors, lib.NewUnaryUUID())
		c.StreamInterceptors = append(c.StreamInterceptors, grpc_opentracing.StreamServerInterceptor())
		c.UnaryInterceptors = append(c.UnaryInterceptors, grpc_opentracing.UnaryServerInterceptor())
		c.StreamInterceptors = append(c.StreamInterceptors, grpc_recovery.StreamServerInterceptor())
		c.UnaryInterceptors = append(c.UnaryInterceptors, grpc_recovery.UnaryServerInterceptor())
		c.StreamInterceptors = append(c.StreamInterceptors, lib.NewStreamLogger())
		c.UnaryInterceptors = append(c.UnaryInterceptors, lib.NewUnaryLogger())
	}
}

func WithDefaultPlugins() Option {
	return func(c *Config) {
		c.Plugins = append(c.Plugins, lib.NewReflection())
		c.Plugins = append(c.Plugins, lib.NewChannelz())
		c.Plugins = append(c.Plugins, lib.NewHealthz())
	}
}
