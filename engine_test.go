package engine_test

import (
	"context"
	"fmt"
	"github.com/autom8ter/engine"
	"github.com/autom8ter/engine/config"
	"github.com/autom8ter/engine/examples/examplepb/client"
	"github.com/autom8ter/util"
	"github.com/grpc-ecosystem/grpc-gateway/examples/proto/examplepb"
	"google.golang.org/grpc/grpclog"
	"os"
	"testing"
	"time"
)

func init() {
	if err := os.Setenv("DEBUG", "t"); err != nil {
		grpclog.Fatalln(err.Error())
	}
}

func TestGRPC(t *testing.T) {
	var eng = engine.New("tcp", ":3002", "Plugin").With(
		config.WithPluginPaths("bin/example.so"),
	)
	go eng.Serve()
	var grpcCli = client.ExampleClient(":3002")
	resp, err := grpcCli.EchoBody(context.Background(), &examplepb.SimpleMessage{
		Id:  "yoyoyoyoyo",
		Num: 199,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println("GRPC RESPONSE:")
	fmt.Println(util.ToPrettyJsonString(resp))
}

func temp() {
	// consider using flags, env vars. or a config file to populate the inputs needed to create an engine instance
	if err := engine.New("tcp", ":3002", "Plugin").With(
		config.WithDebug(),
		config.WithStatsHandler(nil),
		config.WithConnTimeout(2*time.Minute),
		config.WithCreds(nil),
		config.WithMaxConcurrentStreams(1000),
		config.WithPluginPaths("bin/example.so"),
	).Serve(); err != nil {
		grpclog.Fatalln(err.Error())
	}
}
