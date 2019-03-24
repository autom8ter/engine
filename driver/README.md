# driver
--
    import "github.com/autom8ter/engine/driver"


## Usage

#### func  IsPlugin

```go
func IsPlugin(i interface{}) bool
```

#### type GRPCFunc

```go
type GRPCFunc func(s *grpc.Server)
```

PluginFunc is a typed function that satisfies the Plugin interface. It uses a
pattern similar to that of http.Handlerfunc Embed a PluginFunc in a struct to
create a Plugin, and then initialize the PluginFunc with the method generated
from your .pb file

#### func  NewGRPCFunc

```go
func NewGRPCFunc(fn func(s *grpc.Server)) GRPCFunc
```
NewPluginFunc is a helper function to create a PluginFunc

#### func (GRPCFunc) RegisterWithServer

```go
func (p GRPCFunc) RegisterWithServer(s *grpc.Server)
```
RegisterWithServer registers a Plugin with a grpc server.

#### type Plugin

```go
type Plugin interface {
	RegisterWithServer(*grpc.Server)
}
```

Plugin is an interface for representing gRPC server implementations. It is
easily satisfied with code generated from the protoc-gen-go grpc tool

#### func  LoadPlugins

```go
func LoadPlugins(symbol string, paths ...string) []Plugin
```
LoadPlugins loads Plugins from paths set with config.WithPluginPaths(...)

#### func  NewPlugin

```go
func NewPlugin(pluginFunc GRPCFunc) Plugin
```
NewPlugin is a helper function to create a Plugin from a PluginFunc
