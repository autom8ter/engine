//+build

package main

import (
	"context"
	"github.com/autom8ter/engine/driver"
	"github.com/grpc-ecosystem/grpc-gateway/examples/proto/examplepb"
	"google.golang.org/grpc"
)

func init() {
	Plugin = NewExample()
}

var Plugin Example

type Example struct {
	driver.GRPCFunc
}

func NewExample() Example {
	e := Example{
	}
	e.GRPCFunc = func(s *grpc.Server) {
		examplepb.RegisterEchoServiceServer(s, e)
	}
	return e
}

func (e Example) Echo(ctx context.Context, r *examplepb.SimpleMessage) (*examplepb.SimpleMessage, error) {
	return &examplepb.SimpleMessage{
		Id:     r.Id,
		Num:    r.Num,
		Code:   r.Code,
		Status: r.Status,
		Ext:    r.Ext,
	}, nil
}

func (e Example) EchoBody(ctx context.Context, r *examplepb.SimpleMessage) (*examplepb.SimpleMessage, error) {
	return &examplepb.SimpleMessage{
		Id:     r.Id,
		Num:    r.Num,
		Code:   r.Code,
		Status: r.Status,
		Ext:    r.Ext,
	}, nil
}

func (e Example) EchoDelete(ctx context.Context, r *examplepb.SimpleMessage) (*examplepb.SimpleMessage, error) {
	return &examplepb.SimpleMessage{
		Id:     r.Id,
		Num:    r.Num,
		Code:   r.Code,
		Status: r.Status,
		Ext:    r.Ext,
	}, nil
}


/*
rpc Echo(SimpleMessage) returns (SimpleMessage) {
		option (google.api.http) = {
			post: "/v1/example/echo/{id}"
			additional_bindings {
				get: "/v1/example/echo/{id}/{num}"
			}
			additional_bindings {
				get: "/v1/example/echo/{id}/{num}/{lang}"
			}
			additional_bindings {
				get: "/v1/example/echo1/{id}/{line_num}/{status.note}"
			}
			additional_bindings {
				get: "/v1/example/echo2/{no.note}"
			}
		};
	}
	// EchoBody method receives a simple message and returns it.
	rpc EchoBody(SimpleMessage) returns (SimpleMessage) {
		option (google.api.http) = {
			post: "/v1/example/echo_body"
			body: "*"
		};
	}
	// EchoDelete method receives a simple message and returns it.
	rpc EchoDelete(SimpleMessage) returns (SimpleMessage) {
		option (google.api.http) = {
			delete: "/v1/example/echo_delete"
		};
	}
*/