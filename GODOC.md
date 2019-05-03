# engine
--
    import "github.com/autom8ter/engine"


## Usage

#### func  Serve

```go
func Serve(addr string, debug bool, plugs ...driver.Plugin) error
```
Serve creates starts a gRPC server without the need to create an engine instance

#### type Engine

```go
type Engine interface {
	With(opts ...config.Option) *Runtime
	Config() *config.Config
	Shutdown(ctx context.Context)
	Serve() error
}
```

Engine is an interface used to describe a server runtime

#### func  Default

```go
func Default(network, addr string, debug bool) Engine
```
New creates a new engine intstance.

#### func  New

```go
func New(network, addr string, debug bool) Engine
```
New creates a new engine intstance.

#### type Runtime

```go
type Runtime struct {
}
```

Runtime is an implementation of the engine API.

#### func (*Runtime) Config

```go
func (e *Runtime) Config() *config.Config
```
Config returns the runtimes current configuration

#### func (*Runtime) Proxy

```go
func (e *Runtime) Proxy(server http.Server) error
```
Serve starts the runtime gRPC server.

#### func (*Runtime) Serve

```go
func (e *Runtime) Serve() error
```
Serve starts the runtime gRPC server.

#### func (*Runtime) Shutdown

```go
func (e *Runtime) Shutdown(ctx context.Context)
```
Shutdown gracefully closes the grpc server.

#### func (*Runtime) With

```go
func (e *Runtime) With(opts ...config.Option) *Runtime
```
With wraps the runtimes config with config options ref:
github.com/autom8ter/engine/config/options.go
