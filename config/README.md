# config
--
    import "github.com/autom8ter/engine/config"


## Usage

#### type Config

```go
type Config struct {
	Network            string `json:"network" validate:"required"`
	Address            string `json:"address" validate:"required"`
	Plugins            []driver.Plugin
	UnaryInterceptors  []grpc.UnaryServerInterceptor
	StreamInterceptors []grpc.StreamServerInterceptor
	Option             []grpc.ServerOption
}
```

Config contains configurations of gRPC and Gateway server. A new instance of
Config is created from your config.yaml|config.json file in your current working
directory Network, Address, and Paths can be set in your config file to set the
Config instance. Otherwise, defaults are set.

#### func  New

```go
func New(network, addr string, debug bool) *Config
```
New creates a config from your config file. If no config file is present, the
resulting Config will have the following defaults: netowork: "tcp" address:
":3000" use the With method to continue to modify the resulting Config object

#### func (*Config) CreateListener

```go
func (c *Config) CreateListener() (net.Listener, error)
```
CreateListener creates a network listener from the network and address config

#### func (*Config) Debug

```go
func (c *Config) Debug() string
```

#### func (*Config) ServerOptions

```go
func (c *Config) ServerOptions() []grpc.ServerOption
```

#### func (*Config) With

```go
func (c *Config) With(opts ...Option) *Config
```
With is used to configure/initialize a Config with custom options

#### type Option

```go
type Option func(*Config)
```

Option configures a gRPC and a gateway server.

#### func  WithChannelz

```go
func WithChannelz() Option
```
WithChannelz adds grpc server channelz to the list of plugins ref:
https://godoc.org/google.golang.org/grpc/channelz/grpc_channelz_v1

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

#### func  WithDefaultMiddlewares

```go
func WithDefaultMiddlewares() Option
```

#### func  WithDefaultPlugins

```go
func WithDefaultPlugins() Option
```

#### func  WithHealthz

```go
func WithHealthz() Option
```
WithHealthz exposes server's health and it must be imported to enable support
for client-side health checks and adds it to plugins. ref:
https://godoc.org/google.golang.org/grpc/health

#### func  WithMaxConcurrentStreams

```go
func WithMaxConcurrentStreams(num uint32) Option
```
WithMaxConcurrentStreams returns a ServerOption that will apply a limit on the
number of concurrent streams to each ServerTransport.

#### func  WithPlugins

```go
func WithPlugins(svrs ...driver.Plugin) Option
```
WithGoPlugins returns an Option that adds hard-coded Plugins(golang) to the
engine runtime as opposed to go/plugins. See driver.Plugin for the interface
definition.

#### func  WithReflection

```go
func WithReflection() Option
```
WithReflection adds grpc server reflection to the list of plugins ref:
https://godoc.org/google.golang.org/grpc/reflection

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

#### func  WithStreamLoggingMiddleware

```go
func WithStreamLoggingMiddleware() Option
```

#### func  WithStreamRecoveryMiddleware

```go
func WithStreamRecoveryMiddleware() Option
```

#### func  WithStreamTraceMiddleware

```go
func WithStreamTraceMiddleware() Option
```

#### func  WithUnaryInterceptors

```go
func WithUnaryInterceptors(interceptors ...grpc.UnaryServerInterceptor) Option
```
WithUnaryInterceptors returns an Option that sets unary interceptor(s) for a
gRPC server.

#### func  WithUnaryLoggingMiddleware

```go
func WithUnaryLoggingMiddleware() Option
```

#### func  WithUnaryRecoveryMiddleware

```go
func WithUnaryRecoveryMiddleware() Option
```

#### func  WithUnaryTraceMiddleware

```go
func WithUnaryTraceMiddleware() Option
```

#### func  WithUnaryUUIDMiddleware

```go
func WithUnaryUUIDMiddleware() Option
```
