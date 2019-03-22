package driver

import (
	"google.golang.org/grpc"
)

// Plugin is an interface for representing gRPC server implementations. It is easily satisfied with code generated from
// the protoc-gen-go grpc tool
type Plugin interface {
	RegisterWithServer(*grpc.Server)
}

// PluginFunc is a typed function that satisfies the Plugin interface. It uses a pattern similar to that of http.Handlerfunc
// Embed a PluginFunc in a struct to create a Plugin, and then initialize the PluginFunc with the method generated from your .pb file
type PluginFunc func(s *grpc.Server)

// RegisterWithServer registers a Plugin with a grpc server.
func (p PluginFunc) RegisterWithServer(s *grpc.Server) {
	p(s)
}
