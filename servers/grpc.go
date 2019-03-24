package servers

import (
	"github.com/autom8ter/engine/config"
	"github.com/autom8ter/engine/servers/driver"
	"github.com/autom8ter/engine/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/channelz/service"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/reflection"
	"net"
)

// GrpcServer wraps grpc.Server setup process.
type GrpcServer struct {
	server *grpc.Server
}

// NewGrpcServer creates a new GrpcServer instance.
func NewGrpcServer(c *config.Config) driver.Server {
	s := grpc.NewServer(c.ServerOptions...)
	util.Debugln("creating grpc server")
	reflection.Register(s)
	util.Debugln("registered server reflection")
	service.RegisterChannelzServiceToServer(s)
	util.Debugln("registered server channelz")
	for i, svr := range c.Plugins {
		svr.RegisterWithServer(s)
		util.Debugf("plugin count: %v\n", i+1)
	}
	return &GrpcServer{
		server: s,
	}
}

// Serve implements Server.Serve for starting the grpc server
func (s *GrpcServer) Serve(lis net.Listener) error {
	grpclog.Infof("gRPC server is starting %s\n", lis.Addr())
	return s.server.Serve(lis)
}

// Shutdown implements Server.Shutdown for gracefully shutting down the grpc server
func (s *GrpcServer) Shutdown() {
	grpclog.Infoln("shutting down grpc server...")
	s.server.GracefulStop()
}
