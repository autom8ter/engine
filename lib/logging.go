package lib

import (
	"context"
	"github.com/autom8ter/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"time"
)

type Output struct {
	Start string `json:"start"`
	Method string `json:"method"`
	Request interface{} `json:"request"`
	Duration string `json:"duration"`
}


func NewUnaryLogger() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		start := time.Now()
		grpclog.Infoln(util.ToPrettyJsonString(&Output{
			Start:    start.String(),
			Method:   info.FullMethod,
			Request:  req,
			Duration: time.Since(start).String(),
		}))
		return handler(ctx, req)
	}
}

func NewStreamLogger() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		start := time.Now()
		grpclog.Infoln(util.ToPrettyJsonString(&Output{
			Start:    start.String(),
			Method:   info.FullMethod,
			Request:  srv,
			Duration: time.Since(start).String(),
		}))
		return handler(srv, ss)
	}
}