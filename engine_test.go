package engine_test

import (
	"context"
	"fmt"
	"github.com/autom8ter/engine"
	"github.com/autom8ter/engine/config"
	"github.com/autom8ter/engine/lib/examplepb/client"
	"github.com/autom8ter/util"
	"github.com/grpc-ecosystem/grpc-gateway/examples/proto/examplepb"
	"github.com/spf13/viper"
	"testing"
)

func TestClient(t *testing.T) {
	var eng = engine.New().With(
		config.WithNetwork("tcp", ":3000"),
		config.WithPluginPaths("bin/example.plugin"),
		config.WithPluginSymbol("Plugin"),
		config.WithEnvPrefix("ENGINE"),

	)
	go eng.Serve()
	var grpcCli = client.ExampleClient(viper.GetString("address"))
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
