package config

import (
	"crypto/tls"
	"github.com/autom8ter/engine/driver"
	"github.com/autom8ter/engine/util"
	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/stats"
	"log"
	"net/http"
	"os"
	"time"
)

// Option configures a http server.
type HTTPOption func(s *http.Server)

// Option configures a http router.
type RouterOption func(r *mux.Router)

// Option configures a gRPC server.
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
func WithPluginPaths(symbol string, paths ...string) Option {
	return func(c *Config) {
		plugs := driver.LoadPlugins(symbol, paths...)
		if plugs == nil {
			grpclog.Fatalf("failed to load plugins from provided symbol and paths: %s %s", symbol, paths)
		}
		c.Plugins = append(c.Plugins, plugs...)
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
		c.ServerOptions = append(c.ServerOptions, grpc.Creds(creds))
	}
}

// WithStatsHandler returns a ServerOption that sets the stats handler for the server.
func WithStatsHandler(h stats.Handler) Option {
	return func(c *Config) {
		c.ServerOptions = append(c.ServerOptions, grpc.StatsHandler(h))
	}
}

// WithStatsHandler ConnectionTimeout returns a ServerOption that sets the timeout for connection establishment (up to and including HTTP/2 handshaking) for all new connections.
// If this is not set, the default is 120 seconds.
func WithConnTimeout(t time.Duration) Option {
	return func(c *Config) {
		c.ServerOptions = append(c.ServerOptions, grpc.ConnectionTimeout(t))
	}
}

// WithMaxConcurrentStreams returns a ServerOption that will apply a limit on the number
// of concurrent streams to each ServerTransport.
func WithMaxConcurrentStreams(num uint32) Option {
	return func(c *Config) {
		c.ServerOptions = append(c.ServerOptions, grpc.MaxConcurrentStreams(num))
	}
}

// WithMaxConcurrentStreams returns a ServerOption that will apply a limit on the number
// of concurrent streams to each ServerTransport.
func WithRouterOptions(opts ...RouterOption) Option {
	return func(c *Config) {
		c.RouterOptions = append(c.RouterOptions, opts...)
	}
}

// WithHTTPTLS adds a tls config to the http server
func WithHTTPTLS(config *tls.Config) Option {
	return func(c *Config) {
		c.HTTPOptions = append(c.HTTPOptions, func(s *http.Server) {
			s.TLSConfig = config
		})
	}
}

// WithHTTPWriteTO sets the http read header timeout
func WithHTTPWriteTO(dur time.Duration) Option {
	return func(c *Config) {
		c.HTTPOptions = append(c.HTTPOptions, func(s *http.Server) {
			s.WriteTimeout = dur
		})
	}
}

// WithHTTPReadTO sets the http read timeout
func WithHTTPReadTO(dur time.Duration) Option {
	return func(c *Config) {
		c.HTTPOptions = append(c.HTTPOptions, func(s *http.Server) {
			s.ReadTimeout = dur
		})
	}
}

// WithHTTPReadHeaderTO sets the http read header timeout
func WithHTTPReadHeaderTO(dur time.Duration) Option {
	return func(c *Config) {
		c.HTTPOptions = append(c.HTTPOptions, func(s *http.Server) {
			s.ReadHeaderTimeout = dur
		})
	}
}

// WithHTTPTLS adds a tls config to the http server
func WithHTTPErrorLog(lg *log.Logger) Option {
	return func(c *Config) {
		c.HTTPOptions = append(c.HTTPOptions, func(s *http.Server) {
			s.ErrorLog = lg
		})
	}
}
