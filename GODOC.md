# engine
--
    import "github.com/autom8ter/engine"


## Usage

#### type Engine

```go
type Engine interface {
	With(opts ...config.Option) *Runtime
	Shutdown()
	Serve(ctx context.Context) error
}
```

Engine is an interface used to describe a server runtime

#### func  New

```go
func New(network, addr string) Engine
```
New creates a new engine intstance.

#### type Runtime

```go
type Runtime struct {
}
```

Runtime is an implementation of the engine API.

#### func (*Runtime) Serve

```go
func (e *Runtime) Serve(ctx context.Context) error
```
Serve starts the runtime gRPC server.

#### func (*Runtime) Shutdown

```go
func (e *Runtime) Shutdown()
```
Shutdown gracefully closes the grpc server.

#### func (*Runtime) With

```go
func (e *Runtime) With(opts ...config.Option) *Runtime
```
With wraps the runtimes config with config options ref:
github.com/autom8ter/engine/config/options.go
