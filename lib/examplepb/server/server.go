package server

import (
	"context"
	"github.com/autom8ter/engine/driver"
	"github.com/grpc-ecosystem/grpc-gateway/examples/proto/examplepb"
	"google.golang.org/grpc"
)

type Example struct {
	driver.PluginFunc
}

func NewExample() *Example {
	e := &Example{}
	e.PluginFunc = func(s *grpc.Server) {
		examplepb.RegisterEchoServiceServer(s, e)
	}
	return e
}

func (e *Example) Echo(ctx context.Context, r *examplepb.SimpleMessage) (*examplepb.SimpleMessage, error) {
	return &examplepb.SimpleMessage{
		Id:     r.Id,
		Num:    r.Num,
		Code:   r.Code,
		Status: r.Status,
		Ext:    r.Ext,
	}, nil
}

func (e *Example) EchoBody(ctx context.Context, r *examplepb.SimpleMessage) (*examplepb.SimpleMessage, error) {
	return &examplepb.SimpleMessage{
		Id:     r.Id,
		Num:    r.Num,
		Code:   r.Code,
		Status: r.Status,
		Ext:    r.Ext,
	}, nil
}

func (e *Example) EchoDelete(ctx context.Context, r *examplepb.SimpleMessage) (*examplepb.SimpleMessage, error) {
	return &examplepb.SimpleMessage{
		Id:     r.Id,
		Num:    r.Num,
		Code:   r.Code,
		Status: r.Status,
		Ext:    r.Ext,
	}, nil
}
