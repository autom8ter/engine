package util

import (
	"github.com/autom8ter/engine/driver"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	channelz "google.golang.org/grpc/channelz/grpc_channelz_v1"
	"google.golang.org/grpc/grpclog"
	"plugin"
)

// Debugf is grpclog.Infof(format, args...) but only executes if debug=true is set in your config or environmental variables
func Debugf(format string, args ...interface{}) {
	if viper.GetBool("debug") {
		grpclog.Infof(format, args...)
	}
}

// Debugln is grpclog.Infoln(args...) but only executes if debug=true is set in your config or environmental variables
func Debugln(args ...interface{}) {
	if viper.GetBool("debug") {
		grpclog.Infoln(args...)
	}
}

// ChannelzClient creates a new grpc channelz client for connecting to a registered channelz server for debugging.
func ChannelzClient(addr string) channelz.ChannelzClient {
	c, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalln(err.Error())
	}
	return channelz.NewChannelzClient(c)
}

func LoadPlugins() []driver.Plugin {
	var plugs = []driver.Plugin{}
	for _, p := range viper.GetStringSlice("paths") {
		Debugf("registered paths: %v\n", viper.GetStringSlice("paths"))
		plug, err := plugin.Open(p)
		if err != nil {
			grpclog.Fatalln(err.Error())
		}
		sym, err := plug.Lookup(viper.GetString("symbol"))
		if err != nil {
			grpclog.Fatalln(err.Error())
		}

		var asPlugin driver.Plugin
		asPlugin, ok := sym.(driver.Plugin)
		if !ok {
			grpclog.Fatalf("provided plugin: %T does not satisfy Plugin interface\n", sym)
		} else {
			Debugf("registered plugin: %T\n", sym)
			plugs = append(plugs, asPlugin)
		}
	}

	return plugs
}
