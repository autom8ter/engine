package servers

import (
	"github.com/autom8ter/engine/config"
	"github.com/autom8ter/engine/servers/driver"
	"github.com/autom8ter/objectify"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"net"
)

var tool = objectify.New()

// GrpcServer wraps grpc.Server setup process.
type GrpcServer struct {
	server *grpc.Server `validate:"required"`
}

// NewGrpcServer creates a new GrpcServer instance.
func NewGrpcServer(c *config.Config) driver.Server {
	s := grpc.NewServer(c.ServerOptions()...)
	for _, svr := range c.Plugins {
		svr.RegisterWithServer(s)
	}

	gserver := &GrpcServer{
		server: s,
	}
	tool.PanicIfNil(gserver)
	if err := tool.Validate(gserver); err != nil {
		panic(errors.New("failed to validate grpc server"))
	}
	return gserver
}

// Serve implements Server.Serve for starting the grpc server
func (s *GrpcServer) Serve(lis net.Listener) error {
	return s.server.Serve(lis)
}

// Shutdown implements Server.Shutdown for gracefully shutting down the grpc server
func (s *GrpcServer) Shutdown() {
	s.server.GracefulStop()
}
