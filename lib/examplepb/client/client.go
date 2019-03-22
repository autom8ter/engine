package client

import (
	"github.com/grpc-ecosystem/grpc-gateway/examples/proto/examplepb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

func ExampleClient(addr string) examplepb.EchoServiceClient {
	c, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalln(err.Error())
	}
	return examplepb.NewEchoServiceClient(c)
}
