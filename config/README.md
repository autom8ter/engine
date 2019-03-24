# config
--
    import "github.com/autom8ter/engine/config"


## Usage

#### type Config

```go
type Config struct {
	Network            string `json:"network"`
	Address            string `json:"address"`
	Plugins            []driver.Plugin
	UnaryInterceptors  []grpc.UnaryServerInterceptor
	StreamInterceptors []grpc.StreamServerInterceptor
	ServerOptions      []grpc.ServerOption
	RouterOptions      []RouterOption
	HTTPOptions        []HTTPOption
}
```

Config contains configurations of gRPC and Gateway server. A new instance of
Config is created from your config.yaml|config.json file in your current working
directory Network, Address, and Paths can be set in your config file to set the
Config instance. Otherwise, defaults are set.

#### func  New

```go
func New(network, addr string) *Config
```
New creates a config from your config file. If no config file is present, the
resulting Config will have the following defaults: netowork: "tcp" address:
":3000" use the With method to continue to modify the resulting Config object

#### func (*Config) Any

```go
func (c *Config) Any() net.Listener
```

#### func (*Config) GRPC

```go
func (c *Config) GRPC() net.Listener
```

#### func (*Config) HTTP

```go
func (c *Config) HTTP() net.Listener
```

#### func (*Config) HTTPHeaderField

```go
func (c *Config) HTTPHeaderField(name, value string) net.Listener
```

#### func (*Config) Serve

```go
func (c *Config) Serve() error
```

#### func (*Config) With

```go
func (c *Config) With(opts ...Option) *Config
```
With is used to configure/initialize a Config with custom options

#### type HTTPOption

```go
type HTTPOption func(s *http.Server)
```

Option configures a http server.

#### type Option

```go
type Option func(*Config)
```

Option configures a gRPC server.

#### func  WithConnTimeout

```go
func WithConnTimeout(t time.Duration) Option
```
WithStatsHandler ConnectionTimeout returns a ServerOption that sets the timeout
for connection establishment (up to and including HTTP/2 handshaking) for all
new connections. If this is not set, the default is 120 seconds.

#### func  WithCreds

```go
func WithCreds(creds credentials.TransportCredentials) Option
```
WithCreds returns a ServerOption that sets credentials for server connections.

#### func  WithDebug

```go
func WithDebug() Option
```
WithDebug sets debug to true if not already set in your config or environmental
variables

#### func  WithGoPlugins

```go
func WithGoPlugins(svrs ...driver.Plugin) Option
```
WithGoPlugins returns an Option that adds hard-coded Plugins(golang) to the
engine runtime as opposed to go/plugins. See driver.Plugin for the interface
definition.

#### func  WithHTTPErrorLog

```go
func WithHTTPErrorLog(lg *log.Logger) Option
```
WithHTTPTLS adds a tls config to the http server

#### func  WithHTTPReadHeaderTO

```go
func WithHTTPReadHeaderTO(dur time.Duration) Option
```
WithHTTPReadHeaderTO sets the http read header timeout

#### func  WithHTTPReadTO

```go
func WithHTTPReadTO(dur time.Duration) Option
```
WithHTTPReadTO sets the http read timeout

#### func  WithHTTPTLS

```go
func WithHTTPTLS(config *tls.Config) Option
```
WithHTTPTLS adds a tls config to the http server

#### func  WithHTTPWriteTO

```go
func WithHTTPWriteTO(dur time.Duration) Option
```
WithHTTPWriteTO sets the http read header timeout

#### func  WithMaxConcurrentStreams

```go
func WithMaxConcurrentStreams(num uint32) Option
```
WithMaxConcurrentStreams returns a ServerOption that will apply a limit on the
number of concurrent streams to each ServerTransport.

#### func  WithPluginPaths

```go
func WithPluginPaths(symbol string, paths ...string) Option
```
WithPluginPaths adds relative filepaths to Plugins to add to the engine runtime
ref: https://golang.org/pkg/plugin/

#### func  WithRouterOptions

```go
func WithRouterOptions(opts ...RouterOption) Option
```
WithMaxConcurrentStreams returns a ServerOption that will apply a limit on the
number of concurrent streams to each ServerTransport.

#### func  WithStatsHandler

```go
func WithStatsHandler(h stats.Handler) Option
```
WithStatsHandler returns a ServerOption that sets the stats handler for the
server.

#### func  WithStreamInterceptors

```go
func WithStreamInterceptors(interceptors ...grpc.StreamServerInterceptor) Option
```
WithGrpcServerStreamInterceptors returns an Option that sets stream
interceptor(s) for a gRPC server.

#### func  WithUnaryInterceptors

```go
func WithUnaryInterceptors(interceptors ...grpc.UnaryServerInterceptor) Option
```
WithUnaryInterceptors returns an Option that sets unary interceptor(s) for a
gRPC server.

#### type RouterOption

```go
type RouterOption func(r *mux.Router)
```

Option configures a http router.
