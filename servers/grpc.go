package servers

import (
	"github.com/autom8ter/engine/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
	"net"
)

//Server provides an interface for starting and stopping the grpc server.
type Server interface {
	Serve(l net.Listener) error
	Shutdown()
}

// GrpcServer wraps grpc.Server setup process.
type GrpcServer struct {
	server *grpc.Server
	*config.Config
}

// NewGrpcServer creates GrpcServer instance.
func NewGrpcServer(c *config.Config) Server {
	s := grpc.NewServer(c.Option...)
	reflection.Register(s)
	for _, svr := range c.Plugins {
		svr.RegisterWithServer(s)
	}
	return &GrpcServer{
		server: s,
		Config: c,
	}
}

// Serve implements Server.Shutdown
func (s *GrpcServer) Serve(l net.Listener) error {
	grpclog.Infof("gRPC server is starting %s", l.Addr())
	return s.server.Serve(l)
}

// Shutdown implements Server.Shutdown
func (s *GrpcServer) Shutdown() {
	grpclog.Infof("shutting down grpc server...")
	s.server.GracefulStop()
}
