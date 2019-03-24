package logging

import (
	"context"
	"github.com/autom8ter/util"
	"github.com/prometheus/common/log"
	"google.golang.org/grpc"
	"time"
)



func NewUnaryLogger() grpc.UnaryServerInterceptor {
	start := time.Now()
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		log.Infoln("start", start.String())
		log.Infoln("method", info.FullMethod)
		log.Infoln("request", util.ToPrettyJsonString(req))
		log.Infoln("duration", time.Since(start).String())
		return resp, err
	}
}
