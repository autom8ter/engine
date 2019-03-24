package lib

import (
	"github.com/autom8ter/engine/driver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthz "google.golang.org/grpc/health/grpc_health_v1"
)

type Healthz struct {
	driver.PluginFunc
}

func NewHealthz() *Healthz {
	return &Healthz{
		PluginFunc: func(s *grpc.Server) {
			h := health.NewServer()
			healthz.RegisterHealthServer(s, h)
		},
	}
}