package util

import (
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	channelz "google.golang.org/grpc/channelz/grpc_channelz_v1"
	"google.golang.org/grpc/grpclog"
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
