package driver

import (
	"google.golang.org/grpc"
)

//Plugin is an interface for representing gRPC server implementations.
type Plugin interface {
	RegisterWithServer(*grpc.Server)
}

type PluginFunc func(s *grpc.Server)

func (p PluginFunc) RegisterWithServer(s *grpc.Server) {
	p(s)
}
