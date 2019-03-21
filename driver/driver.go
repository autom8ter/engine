package driver

import (
	"google.golang.org/grpc"
	"net"
)

// Handler is an interface for representing gRPC server implementations.
type Plugin interface {
	RegisterWithServer(*grpc.Server)
}

// Server provides an interface for starting and stopping the server.
type Server interface {
	Serve(l net.Listener) error
	Shutdown()
}
