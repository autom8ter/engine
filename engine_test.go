package engine_test

import (
	"context"
	"fmt"
	"github.com/autom8ter/engine/lib/examplepb/client"
	"github.com/autom8ter/util"
	"github.com/grpc-ecosystem/grpc-gateway/examples/proto/examplepb"
	"github.com/spf13/viper"
	"testing"
)

/*
var eng = engine.New().With(
	config.WithDefaultLogger(),
	config.WithAddr("tcp", ":3000"),
	config.WithPlugins(),
	config.WithRouterWare(
		handlers.DebugWare(),
		handlers.MetricsWare(),
	),
)
*/
var addr = viper.GetString("address")

var grpcCli = client.ExampleClient(addr)

func TestNewEngine(t *testing.T) {
	resp, err := grpcCli.EchoBody(context.Background(), &examplepb.SimpleMessage{
		Id:  "yoyoyoyoyo",
		Num: 199,
	})
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println("GRPC RESPONSE")
	fmt.Println(util.ToPrettyJsonString(resp))
}
