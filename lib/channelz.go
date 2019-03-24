package lib

import (
	"github.com/autom8ter/engine/driver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/channelz/service"
	"google.golang.org/grpc/grpclog"
	channelz "google.golang.org/grpc/channelz/grpc_channelz_v1"
)

type Channelz struct {
	driver.PluginFunc
}

func NewChannelz() *Channelz {
	return &Channelz{
		PluginFunc: func(s *grpc.Server) {
			service.RegisterChannelzServiceToServer(s)
		},
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
