package driver

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

// Handler is an interface for representing gRPC server implementations.
type Plugin interface {
	RegisterWithServer(*grpc.Server)
	RegisterWithHandler(context.Context, *runtime.ServeMux, *grpc.ClientConn) error
}

// Server provides an interface for starting and stopping the server.
type Listener interface {
	Serve(l net.Listener) error
	Shutdown()
}

// PassedHeaderDeciderFunc returns true if given header should be passed to gRPC server metadata.
type PassedHeaderDeciderFunc func(string) bool

// HTTPServerMiddleware is an interface of http server middleware
type HTTPServerMiddleware func(http.Handler) http.Handler
