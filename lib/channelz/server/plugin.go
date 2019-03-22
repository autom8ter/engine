package server

import (
	"github.com/autom8ter/engine/driver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/channelz/service"
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
