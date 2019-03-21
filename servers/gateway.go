package servers

import (
	"context"
	"github.com/autom8ter/engine/config"
	"github.com/autom8ter/engine/driver"
	"github.com/gorilla/mux"
	"net"
	"net/http"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

// NewGatewayServer creates GrpcServer instance.
func NewGatewayServer(c *config.Config) driver.Server {
	return &GatewayServer{
		Config: c,
	}
}

// GatewayListener wraps gRPC gateway server setup process.
type GatewayServer struct {
	server *http.Server
	*config.Config
}

// Serve implements Server.Shutdown
func (s *GatewayServer) Serve(l net.Listener) error {
	conn, err := s.createConn()
	if err != nil {
		return errors.Wrap(err, "failed to create connection with grpc-gateway server")
	}
	defer conn.Close()

	s.server, err = s.createServer(conn)
	if err != nil {
		return errors.Wrap(err, "failed to create grpc-gateway server: %v")
	}

	grpclog.Infof("grpc-gateway server is starting %s", l.Addr())

	return s.server.Serve(l)
}

// Shutdown implements Server.Shutdown
func (s *GatewayServer) Shutdown() {
	grpclog.Infof("shutting down grpc-gateway server...")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err := s.server.Shutdown(ctx)
	if err != nil {
		grpclog.Errorf("failed to shutdown grpc-gateway server: %v", err)
	}
}

func (s *GatewayServer) createConn() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(s.GrpcInternalAddr.Addr, s.ClientOptions()...)
	if err != nil {
		err = errors.Wrap(err, "failed to connect to gRPC server from grpc-gateway server")
	}
	return conn, err
}

func (s *GatewayServer) createServer(conn *grpc.ClientConn) (*http.Server, error) {
	grpclog.Infof("creating grpc-gateway serve mux...")
	rmux := runtime.NewServeMux(
		append(
			[]runtime.ServeMuxOption{runtime.WithProtoErrorHandler(runtime.DefaultHTTPProtoErrorHandler)},
			s.GatewayMuxOptions...,
		)...,
	)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	for _, svr := range s.Plugins {
		err := svr.RegisterWithHandler(ctx, rmux, conn)
		if err != nil {
			return nil, errors.Wrap(err, "failed to register handler for grpc-gateway server")
		}
	}

	var handler http.Handler = rmux

	for i := len(s.GatewayServerMiddlewares) - 1; i >= 0; i-- {
		grpclog.Infof("registering grpc-gateway middleware...")

		handler = (s.GatewayServerMiddlewares[i])(handler)
	}
	router := mux.NewRouter()

	for _, r := range s.RouterWare {
		grpclog.Infof("registering grpc-gateway routerware...")
		router = r(router)
	}
	grpclog.Infof("configuring grpc-gateway server...")
	svr := &http.Server{
		Handler: handler,
	}
	if cfg := s.GatewayServerConfig; cfg != nil {
		cfg.ApplyTo(svr)
	}

	return svr, nil
}
