package config

import (
	"github.com/autom8ter/engine/driver"
	"github.com/autom8ter/engine/util"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/stats"
	"os"
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

// WithPluginPaths adds relative filepaths to Plugins to add to the engine runtime
//ref: https://golang.org/pkg/plugin/
func WithPlugins(symbol string, paths ...string) Option {
	return func(c *Config) {
		c.Symbol = symbol
		c.Paths = append(c.Paths, paths...)
		c.loadPlugins()
	}
}

// WithGoPlugins returns an Option that adds hard-coded Plugins(golang) to the engine runtime as opposed to go/plugins.
// See driver.Plugin for the interface definition.
func WithGoPlugins(svrs ...driver.Plugin) Option {
	return func(c *Config) {
		c.Plugins = append(c.Plugins, svrs...)
		if len(c.Plugins) == 0 {
			grpclog.Fatal(errors.New("zero valid plugins registered"))
		}
	}
}

// WithDebug sets debug to true if not already set in your config or environmental variables
func WithDebug() Option {
	return func(c *Config) {
		util.Debugln("enabling debug mode")
		if err := os.Setenv("DEBUG", "t"); err != nil {
			grpclog.Fatalln(err.Error())
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
