package servers

import (
	"github.com/autom8ter/engine/driver"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc/grpclog"
	"net"
)

// MuxServer wraps a connection multiplexer and a listener.
type MuxServer struct {
	mux cmux.CMux
	lis net.Listener
}

// NewMuxServer creates MuxServer instance.
func NewMuxServer(mux cmux.CMux, lis net.Listener) driver.Server {
	return &MuxServer{
		mux: mux,
		lis: lis,
	}
}

// Serve implements Server.Serve
func (s *MuxServer) Serve(net.Listener) error {
	grpclog.Info("mux-server is starting %s", s.lis.Addr())

	return s.mux.Serve()
}

// Shutdown implements Server.Shutdown
func (s *MuxServer) Shutdown() {
	grpclog.Infof("shutting down mux-server...")
	err := s.lis.Close()
	if err != nil {
		grpclog.Errorf("failed to close mux-server's listener: %v", err)
	}
}
