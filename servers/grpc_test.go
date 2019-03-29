package servers_test

import (
	"github.com/autom8ter/engine/config"
	"github.com/autom8ter/engine/servers"
	"google.golang.org/grpc/grpclog"
	"log"
	"os"
	"testing"
)

func init() {
	if err := os.Setenv("DEBUG", "t"); err != nil {
		grpclog.Fatalln(err.Error())
	}
}

var c = config.New("tcp", ":3005", true)

func TestGrpcServer_Serve(t *testing.T) {
	var s = servers.NewGrpcServer(c)
	lis, err := c.CreateListener()
	if err != nil {
		grpclog.Fatalln(err.Error())
	}

	go s.Serve(lis)
}

func TestNewGrpcServer(t *testing.T) {
	var s = servers.NewGrpcServer(c)
	if s == nil {
		log.Fatal("nil server")
	}
}

func TestGrpcServer_Shutdown(t *testing.T) {
	var s = servers.NewGrpcServer(c)
	if s == nil {
		log.Fatal("nil server")
	}
	s.Shutdown()
}
