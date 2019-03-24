package driver

import (
	"github.com/autom8ter/engine/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

// LoadPlugins loads Plugins from paths set with config.WithPluginPaths(...)
func LoadPlugins(symbol string, paths ...string) []Plugin {
	var plug = []Plugin{}
	for _, p := range paths {
		util.Debugf("registered path: %v\n", p)
		var asPlugin Plugin
		sym := util.Load(p, symbol)
		asPlugin, ok := util.Load(p, symbol).(Plugin)
		if !ok {
			grpclog.Fatalf("provided plugin: %T does not satisfy Plugin interface\n", sym)
		} else {
			util.Debugf("registered plugin: %T\n", sym)
			plug = append(plug, asPlugin)
		}
	}
	return plug
}

func IsPlugin(i interface{}) bool {
	_, ok := i.(Plugin)
	return ok
}

// Plugin is an interface for representing gRPC server implementations. It is easily satisfied with code generated from
// the protoc-gen-go grpc tool
type Plugin interface {
	RegisterWithServer(*grpc.Server)
}

// NewPlugin is a helper function to create a Plugin from a PluginFunc
func NewPlugin(pluginFunc GRPCFunc) Plugin {
	return pluginFunc
}

// PluginFunc is a typed function that satisfies the Plugin interface. It uses a pattern similar to that of http.Handlerfunc
// Embed a PluginFunc in a struct to create a Plugin, and then initialize the PluginFunc with the method generated from your .pb file
type GRPCFunc func(s *grpc.Server)

// NewPluginFunc is a helper function to create a PluginFunc
func NewGRPCFunc(fn func(s *grpc.Server)) GRPCFunc {
	return fn
}

// RegisterWithServer registers a Plugin with a grpc server.
func (p GRPCFunc) RegisterWithServer(s *grpc.Server) {
	p(s)
}
