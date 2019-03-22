package driver

import "net"

// Server provides an interface for starting and stopping the grpc server.
type Server interface {
	Serve(l net.Listener) error
	Shutdown()
}
