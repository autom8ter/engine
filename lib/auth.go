package lib

import (
	"context"
	"github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"google.golang.org/grpc"
)

func UnaryAuthInterceptor(ctxFunc func(ctx context.Context) (context.Context, error)) grpc.UnaryServerInterceptor {
	return grpc_auth.UnaryServerInterceptor(ctxFunc)
}
