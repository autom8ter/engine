package driver

import "net"

type Server interface {
	Shutdown()
	Serve(lis net.Listener) error
}
