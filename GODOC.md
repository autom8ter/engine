# engine
--
    import "github.com/autom8ter/engine"


## Usage

#### type Engine

```go
type Engine interface {
	http.Handler
	With(opts ...config.Option) *Runtime
	Config() *config.Config
	ServeGRPC() error
	Shutdown()
}
```


#### func  New

```go
func New() Engine
```
New creates a engine intstance.

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
Config returns the runtimes current configuration

#### func (*Runtime) ServeGRPC

```go
func (e *Runtime) ServeGRPC() error
```
Serve starts the runtime gRPC server.

#### func (*Runtime) ServeHTTP

```go
func (e *Runtime) ServeHTTP(w http.ResponseWriter, r *http.Request)
```

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
