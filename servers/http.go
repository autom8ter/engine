package servers

import (
	"context"
	"github.com/autom8ter/engine/config"
	"github.com/autom8ter/engine/servers/driver"
	"github.com/autom8ter/engine/util"
	"github.com/gorilla/mux"
	"google.golang.org/grpc/grpclog"
	"net"
	"net/http"
)

type HTTPServer struct {
	server *http.Server
}

func NewHTTPServer(c *config.Config) driver.Server {
	r := mux.NewRouter()
	for _, o := range c.RouterOptions {
		o(r)
	}

	s := &http.Server{
		Addr:    c.Address,
		Handler: r,
	}
	for _, o := range c.HTTPOptions {
		o(s)
	}
	return &HTTPServer{
		server: s,
	}
}

func (s *HTTPServer) Shutdown() {
	util.Debugln("shutting down http server")
	grpclog.Warningln(s.server.Shutdown(context.TODO()))
}

func (s *HTTPServer) Serve(lis net.Listener) error {
	return s.server.Serve(lis)
}
