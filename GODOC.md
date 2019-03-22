# engine
--
    import "github.com/autom8ter/engine"


## Usage

#### type Engine

```go
type Engine interface {
	With(opts ...config.Option) *Runtime
	Config() *config.Config
	Serve() error
	Shutdown()
}
```


#### func  New

```go
func New() Engine
```
New creates a server intstance.

#### type Runtime

```go
type Runtime struct {
}
```

Engine is the framework instance.

#### func (*Runtime) Config

```go
func (e *Runtime) Config() *config.Config
```
New creates a server intstance.

#### func (*Runtime) Serve

```go
func (e *Runtime) Serve() error
```
Serve starts gRPC and Gateway servers.

#### func (*Runtime) Shutdown

```go
func (e *Runtime) Shutdown()
```
Shutdown closes servers.

#### func (*Runtime) With

```go
func (e *Runtime) With(opts ...config.Option) *Runtime
```
New creates a server intstance.
