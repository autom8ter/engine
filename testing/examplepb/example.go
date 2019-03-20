package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/examples/proto/examplepb"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
)

var ExamplePb  Example

type Example struct {
}

func NewExample() *Example {
	return &Example{}
}

func (e *Example) Echo(ctx context.Context, r *examplepb.SimpleMessage) (*examplepb.SimpleMessage, error) {
	return &examplepb.SimpleMessage{
		Id:                   r.Id,
		Num:                  r.Num,
		Code:                 r.Code,
		Status:               r.Status,
		Ext:                  r.Ext,
	}, nil
}

func (e *Example) EchoBody(ctx context.Context, r *examplepb.SimpleMessage) (*examplepb.SimpleMessage, error) {
	return &examplepb.SimpleMessage{
		Id:                   r.Id,
		Num:                  r.Num,
		Code:                 r.Code,
		Status:               r.Status,
		Ext:                  r.Ext,
	}, nil
}

func (e *Example) EchoDelete(ctx context.Context, r *examplepb.SimpleMessage) (*examplepb.SimpleMessage, error) {
	return &examplepb.SimpleMessage{
		Id:                   r.Id,
		Num:                  r.Num,
		Code:                 r.Code,
		Status:               r.Status,
		Ext:                  r.Ext,
	}, nil
}

func (e *Example) RegisterWithServer(s *grpc.Server) {

	examplepb.RegisterEchoServiceServer(s, e)
}

func (e *Example) RegisterWithHandler(ctx context.Context, m *runtime.ServeMux, cc *grpc.ClientConn) error {
	return examplepb.RegisterEchoServiceHandler(ctx, m, cc)
}
