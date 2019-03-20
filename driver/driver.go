package driver

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
	"net"
)

type Wrapper interface {
	Wrap(router *mux.Router) *mux.Router
}

// Handler is an interface for representing gRPC server implementations.
type Plugin interface {
	RegisterWithServer(*grpc.Server)
	RegisterWithHandler(context.Context, *runtime.ServeMux, *grpc.ClientConn) error
}

type RunFunc func(cmd *cobra.Command, args []string)

type Loader interface {
	LoadPlugin(name string) (Plugin, error)
	LoadWrapper(name string) (Wrapper, error)
}

// Server provides an interface for starting and stopping the server.
type Engine interface {
	Serve(l net.Listener) error
	Shutdown()
}
