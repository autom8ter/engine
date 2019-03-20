package engine

import (
	"github.com/autom8ter/engine/config"
	"github.com/autom8ter/engine/driver"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
	"net"
)

// GrpcServer wraps grpc.Server setup process.
type GrpcServer struct {
	server *grpc.Server
	*config.Config
}

// NewGrpcServer creates GrpcServer instance.
func NewGrpcServer(c *config.Config) driver.Engine {
	s := grpc.NewServer(c.ServerOptions()...)
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

	err := s.server.Serve(l)

	grpclog.Infof("gRPC server stopped: %v", err)

	return errors.Wrap(err, "failed to serve gRPC server")
}

// Shutdown implements Server.Shutdown
func (s *GrpcServer) Shutdown() {
	s.server.GracefulStop()
}
