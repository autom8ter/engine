package lib

import (
	"github.com/autom8ter/engine/driver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Reflection struct {
	driver.PluginFunc
}

func NewReflection() *Channelz {
	return &Channelz{
		PluginFunc: func(s *grpc.Server) {
			reflection.Register(s)
		},
	}
}
