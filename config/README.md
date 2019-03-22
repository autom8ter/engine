# config
--
    import "github.com/autom8ter/engine/config"


## Usage

#### type Config

```go
type Config struct {
	Network            string   `mapstructure:"network" json:"network"`
	Address            string   `mapstructure:"address" json:"address"`
	Paths              []string `mapstructure:"paths" json:"paths"`
	Symbol             string   `mapstructure:"symbol" json:"symbol"`
	Plugins            []driver.Plugin
	UnaryInterceptors  []grpc.UnaryServerInterceptor
	StreamInterceptors []grpc.StreamServerInterceptor
	Option             []grpc.ServerOption
}
```


#### func  New

```go
func New() *Config
```
New creates a config from your config file. If no config file is present, the
resulting Config will have the following defaults: netowork: "tcp" address:
":3000" use the With method to continue to modify the resulting Config object

#### func (*Config) CreateListener

```go
func (c *Config) CreateListener() (net.Listener, error)
```
CreateListener creates a network listener for the grpc server from the netowork
address

#### func (*Config) With

```go
func (c *Config) With(opts []Option) *Config
```
With is used to configure/initialize a Config with custom options

#### type Option

```go
type Option func(*Config)
```

Option configures a gRPC and a gateway server.

#### func  WithEnvPrefix

```go
func WithEnvPrefix(prefix string) Option
```
WithGoPlugins returns an Option that adds hard-coded Plugins(golang) to the
engine runtime as opposed to go/plugins.

#### func  WithGoPlugins

```go
func WithGoPlugins(svrs ...driver.Plugin) Option
```
WithGoPlugins returns an Option that adds hard-coded Plugins(golang) to the
engine runtime as opposed to go/plugins.

#### func  WithNetwork

```go
func WithNetwork(network, addr string) Option
```
WithNetwork returns an Option that sets an network address for a gRPC server.

#### func  WithPluginPaths

```go
func WithPluginPaths(paths ...string) Option
```
WithPluginPaths adds relative filepaths to Plugins to add to the engine runtime

#### func  WithPluginSymbol

```go
func WithPluginSymbol(sym string) Option
```
WithGoPlugins returns an Option that adds hard-coded Plugins(golang) to the
engine runtime as opposed to go/plugins.

#### func  WithServerOptions

```go
func WithServerOptions(opts ...grpc.ServerOption) Option
```
WithOptions returns an Option that sets grpc.ServerOption(s) to a gRPC server.

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
