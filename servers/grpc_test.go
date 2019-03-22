package servers_test

import (
	"github.com/autom8ter/engine/config"
	"github.com/autom8ter/engine/servers"
	"log"
	"testing"
)

func TestGrpcServer_Serve(t *testing.T) {
	c := config.New()
	var s = servers.NewGrpcServer(c)
	if s == nil {
		log.Fatal("nil server")
	}
	lis, err := c.CreateListener()
	if err != nil {
		t.Fatal(err.Error())
	}

	go s.Serve(lis)
}

func TestNewGrpcServer(t *testing.T) {
	var s = servers.NewGrpcServer(config.New())
	if s == nil {
		log.Fatal("nil server")
	}
}

func TestGrpcServer_Shutdown(t *testing.T) {
	var s = servers.NewGrpcServer(config.New())
	if s == nil {
		log.Fatal("nil server")
	}
	s.Shutdown()
}
