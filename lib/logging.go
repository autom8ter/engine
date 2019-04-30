package lib

import (
	"context"
	"github.com/autom8ter/objectify"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"time"
)

var tool = objectify.Default()

type output struct {
	UUID       string      `json:"uuid,omitempty"`
	Start      string      `json:"start,omitempty"`
	HumanStart string      `json:"human_start,omitempty"`
	Method     string      `json:"method,omitempty"`
	Request    interface{} `json:"request,omitempty"`
	Duration   string      `json:"duration,omitempty"`
}

type streamOutput struct {
	UUID       string      `json:"uuid,omitempty"`
	Start      string      `json:"start,omitempty"`
	HumanStart string      `json:"human_start,omitempty"`
	Method     string      `json:"method,omitempty"`
	IsClient   bool        `json:"is_client,omitempty"`
	IsServer   bool        `json:"is_server,omitempty"`
	Request    interface{} `json:"request,omitempty"`
	Duration   string      `json:"duration,omitempty"`
}

func NewUnaryLogger() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		start := time.Now()

		grpclog.Infoln(string(tool.MarshalJSON(&output{
			UUID:       uUIDFromContext(ctx),
			Start:      start.String(),
			HumanStart: tool.HumanizeTime(start),
			Method:     info.FullMethod,
			Request:    string(tool.MarshalJSON(req)),
			Duration:   time.Since(start).String(),
		})))
		return handler(ctx, req)
	}
}

func NewStreamLogger() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		start := time.Now()
		grpclog.Infoln(string(tool.MarshalJSON(&streamOutput{
			UUID:       uUIDFromContext(ss.Context()),
			HumanStart: tool.HumanizeTime(start),
			IsClient:   info.IsClientStream,
			IsServer:   info.IsServerStream,
			Start:      start.String(),
			Method:     info.FullMethod,
			Request:    srv,
			Duration:   time.Since(start).String(),
		})))
		return handler(srv, ss)
	}
}
