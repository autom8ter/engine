package main

import (
	"context"
	"fmt"
	"github.com/autom8ter/util"
	"github.com/grpc-ecosystem/grpc-gateway/examples/proto/examplepb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func main() {
	c, err := grpc.DialContext(
		context.TODO(),
		"localhost:8080",
		grpc.WithInsecure(),
	)
	if err != nil {
		grpclog.Fatalln(err.Error())
	}
	client := examplepb.NewEchoServiceClient(c)
	resp, err := client.Echo(context.Background(), &examplepb.SimpleMessage{
		Id: "this is my message",
	})
	if err != nil {
		grpclog.Fatalln(err.Error())
	}
	fmt.Println(util.ToPrettyJsonString(resp))
}
