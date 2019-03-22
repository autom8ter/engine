# driver
--
    import "github.com/autom8ter/engine/driver"


## Usage

#### type Plugin

```go
type Plugin interface {
	RegisterWithServer(*grpc.Server)
}
```

Plugin is an interface for representing gRPC server implementations.

#### type PluginFunc

```go
type PluginFunc func(s *grpc.Server)
```


#### func (PluginFunc) RegisterWithServer

```go
func (p PluginFunc) RegisterWithServer(s *grpc.Server)
```
