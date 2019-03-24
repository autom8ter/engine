package main

import (
	"context"
	"fmt"
	"github.com/autom8ter/engine/examples/examplepb/client"
	"github.com/autom8ter/util"
	"github.com/grpc-ecosystem/grpc-gateway/examples/proto/examplepb"
	"log"
)

func main() {
	c := client.ExampleClient("localhost:8080")
	resp, err := c.Echo(context.Background(), &examplepb.SimpleMessage{
		Id: "this is my message",
	})
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(util.ToPrettyJsonString(resp))
}
