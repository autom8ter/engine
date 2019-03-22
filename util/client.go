package util

import (
	"google.golang.org/grpc"
	channelz "google.golang.org/grpc/channelz/grpc_channelz_v1"
	"google.golang.org/grpc/grpclog"
)

// ChannelzClient creates a new grpc channelz client for connecting to a registered channelz server for debugging.
func ChannelzClient(addr string) channelz.ChannelzClient {
	c, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalln(err.Error())
	}
	return channelz.NewChannelzClient(c)
}
