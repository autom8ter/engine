package driver

import "net"

// Server provides an interface for starting and stopping the grpc server.
type Server interface {
	Serve(l net.Listener) error
	Shutdown()
}

func IsServer(a interface{}) bool {
	_, ok := a.(Server)
	return ok
}
