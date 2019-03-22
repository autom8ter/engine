package servers

import (
	"github.com/autom8ter/engine/config"
	"github.com/autom8ter/engine/util"
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
	util.Debugln("creating grpc server")
	reflection.Register(s)
	util.Debugln("registered server reflection")
	for i, svr := range c.Plugins {
		svr.RegisterWithServer(s)
		util.Debugf("plugin count: %v\n", i+1)
	}
	return &GrpcServer{
		server: s,
		Config: c,
	}
}

// Serve implements Server.Shutdown
func (s *GrpcServer) Serve(l net.Listener) error {
	grpclog.Infof("gRPC server is starting %s\n", l.Addr())
	return s.server.Serve(l)
}

// Shutdown implements Server.Shutdown
func (s *GrpcServer) Shutdown() {
	grpclog.Infoln("shutting down grpc server...")
	s.server.GracefulStop()
}
