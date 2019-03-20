package helloworld

import (
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"strings"
)

type GreeterFunc func(ctx context.Context, req *HelloRequest) (*HelloReply, error)

type Greeter struct {
	Func GreeterFunc
}

func NewGreeter(fn GreeterFunc) *Greeter {
	return &Greeter{
		fn,
	}
}

func NewBasicGreeter() *Greeter {
	return &Greeter{
		Func: func(ctx context.Context, req *HelloRequest) (reply *HelloReply, e error) {
			reply.Message = fmt.Sprintf("Hello %s!", strings.ToTitle(req.Name))
			return
		},
	}
}

func (g *Greeter) SayHello(ctx context.Context, req *HelloRequest) (*HelloReply, error) {
	return g.Func(ctx, req)
}

func (g *Greeter) RegisterWithServer(s *grpc.Server) {
	RegisterGreeterServer(s, g)
}

func (g *Greeter) RegisterWithHandler(ctx context.Context, m *runtime.ServeMux, cc *grpc.ClientConn) error {
	return RegisterGreeterHandler(ctx, m, cc)
}
