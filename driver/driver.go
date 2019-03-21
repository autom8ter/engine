/*
Example plugin:


*/

package driver

import (
	"google.golang.org/grpc"
	"net"
)

// Plugin is an interface for representing gRPC server implementations.
type Plugin interface {
	RegisterWithServer(*grpc.Server)
}

// Server provides an interface for starting and stopping the grpc server.
type Server interface {
	Serve(l net.Listener) error
	Shutdown()
}

//PluginFunc implements the Plugin interface.
type PluginFunc func(*grpc.Server)

func (p PluginFunc) RegisterWithServer(s *grpc.Server) {
	p(s)
}
