package driver_test

import (
	"github.com/autom8ter/engine/config"
	"github.com/autom8ter/engine/servers"
	"github.com/autom8ter/engine/servers/driver"
	"testing"
)

func TestIsServer(t *testing.T) {
	s := servers.NewGrpcServer(config.New())
	if !driver.IsServer(s) {
		t.Fatal("not a server")
	}
}
