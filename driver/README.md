# driver
--
    import "github.com/autom8ter/engine/driver"


## Usage

#### func  IsPlugin

```go
func IsPlugin(a interface{}) bool
```

#### type Plugin

```go
type Plugin interface {
	RegisterWithServer(*grpc.Server)
}
```

Plugin is an interface for representing gRPC server implementations. It is
easily satisfied with code generated from the protoc-gen-go grpc tool

#### func  NewPlugin

```go
func NewPlugin(pluginFunc PluginFunc) Plugin
```
NewPlugin is a helper function to create a Plugin from a PluginFunc

#### type PluginFunc

```go
type PluginFunc func(s *grpc.Server)
```

PluginFunc is a typed function that satisfies the Plugin interface. It uses a
pattern similar to that of http.Handlerfunc Embed a PluginFunc in a struct to
create a Plugin, and then initialize the PluginFunc with the method generated
from your .pb file

#### func  NewPluginFunc

```go
func NewPluginFunc(fn func(s *grpc.Server)) PluginFunc
```
NewPluginFunc is a helper function to create a PluginFunc

#### func (PluginFunc) RegisterWithServer

```go
func (p PluginFunc) RegisterWithServer(s *grpc.Server)
```
RegisterWithServer registers a Plugin with a grpc server.
