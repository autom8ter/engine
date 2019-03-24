package logging

import (
	"context"
	"github.com/autom8ter/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"time"
)

type Output struct {
	Time     time.Time     `json:"time"`
	Method   string        `json:"method"`
	Request  interface{}   `json:"request"`
	Duration time.Duration `json:"duration"`
}

func NewOutput(time time.Time, method string, request interface{}, duration time.Duration) *Output {
	return &Output{Time: time, Method: method, Request: request, Duration: duration}
}

func (o *Output) Error() string {
	return util.ToPrettyJsonString(o)
}

func NewUnaryLogger() grpc.UnaryServerInterceptor {
	start := time.Now()
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		grpclog.Infoln(NewOutput(start, info.FullMethod, req, time.Since(start)).Error())
		return resp, err
	}
}
