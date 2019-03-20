package examplepb

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/examples/proto/examplepb"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

type Example struct {
}

func NewExample() *Example {
	return &Example{}
}

func (e *Example) Echo(context.Context, *examplepb.SimpleMessage) (*examplepb.SimpleMessage, error) {
	panic("implement me")
}

func (e *Example) EchoBody(context.Context, *examplepb.SimpleMessage) (*examplepb.SimpleMessage, error) {
	panic("implement me")
}

func (e *Example) EchoDelete(context.Context, *examplepb.SimpleMessage) (*examplepb.SimpleMessage, error) {
	panic("implement me")
}

func (e *Example) RegisterWithServer(s *grpc.Server) {

	examplepb.RegisterEchoServiceServer(s, e)
}

func (e *Example) RegisterWithHandler(ctx context.Context, m *runtime.ServeMux, cc *grpc.ClientConn) error {
	return examplepb.RegisterEchoServiceHandler(ctx, m, cc)
}
