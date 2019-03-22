# servers
--
    import "github.com/autom8ter/engine/servers"


## Usage

#### type GrpcServer

```go
type GrpcServer struct {
	*config.Config
}
```

GrpcServer wraps grpc.Server setup process.

#### func (*GrpcServer) Serve

```go
func (s *GrpcServer) Serve(l net.Listener) error
```
Serve implements Server.Shutdown

#### func (*GrpcServer) Shutdown

```go
func (s *GrpcServer) Shutdown()
```
Shutdown implements Server.Shutdown

#### type Server

```go
type Server interface {
	Serve(l net.Listener) error
	Shutdown()
}
```

Server provides an interface for starting and stopping the grpc server.

#### func  NewGrpcServer

```go
func NewGrpcServer(c *config.Config) Server
```
NewGrpcServer creates GrpcServer instance.
