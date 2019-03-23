package driver

import (
	"net"
	"net/http"
)

// Server provides an interface for starting and stopping the grpc server.
type Server interface {
	http.Handler
	Serve(l net.Listener) error
	Shutdown()
}

func IsServer(a interface{}) bool {
	_, ok := a.(Server)
	return ok
}
