/*
Example plugin:

```
var Plugin  Example

type Example struct {
	driver.PluginFunc
}

func NewExample() *Example {
	e := &Example{}
	e.PluginFunc = func(s *grpc.Server) {
		examplepb.RegisterEchoServiceServer(s, e)
	}
	return e
}
```

*/

package driver

import (
	"google.golang.org/grpc"
	"net"
)

//Plugin is an interface for representing gRPC server implementations.
type Plugin interface {
	RegisterWithServer(*grpc.Server)
}

//Server provides an interface for starting and stopping the grpc server.
type Server interface {
	Serve(l net.Listener) error
	Shutdown()
}

//PluginFunc implements the Plugin interface.
type PluginFunc func(*grpc.Server)

//RegisterWithServer is an interface for representing gRPC server implementations.
func (p PluginFunc) RegisterWithServer(s *grpc.Server) {
	p(s)
}
