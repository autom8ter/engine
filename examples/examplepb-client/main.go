package main

import (
	"context"
	"fmt"
	"github.com/autom8ter/util"
	"github.com/grpc-ecosystem/grpc-gateway/examples/proto/examplepb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"log"
)

func InsecureClient(addr string) examplepb.EchoServiceClient {
	c, err := grpc.Dial(
		addr,
		grpc.WithInsecure(),
	)
	if err != nil {
		grpclog.Fatalln(err.Error())
	}
	return examplepb.NewEchoServiceClient(c)
}

func SecureClient(addr string) examplepb.EchoServiceClient {
	ctx := context.WithValue(context.TODO(), "bearer", "i-come-in-peace")
	c, err := grpc.DialContext(
		ctx,
		addr,
		grpc.WithInsecure(),
	)
	if err != nil {
		grpclog.Fatalln(err.Error())
	}
	return examplepb.NewEchoServiceClient(c)
}

func main() {
	c := InsecureClient("localhost:8080")
	resp, err := c.Echo(context.Background(), &examplepb.SimpleMessage{
		Id: "this is my message",
	})
	if err != nil {
		grpclog.Warningln(err.Error())
	}
	log.Println("Insecure Response")
	fmt.Println(util.ToPrettyJsonString(resp))
	c = SecureClient("localhost:8080")
	resp, err = c.Echo(context.Background(), &examplepb.SimpleMessage{
		Id: "this is my message",
	})
	if err != nil {
		grpclog.Warningln(err.Error())
	}
	log.Println("Secure Response")
	fmt.Println(util.ToPrettyJsonString(resp))
}
