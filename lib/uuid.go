package lib

import (
	"context"
	"github.com/autom8ter/engine/util"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
)

type keyType int

const RequestIDKey = keyType(iota)

// returns a string that can be used as a request ID for client requests that don't already have one.
func newUUID() (string, error) {
	v, err := uuid.NewUUID()
	if err != nil {
		return "", err
	}
	return v.String(), nil
}

// uUIDFromContext return the request ID found on the supplied context
func uUIDFromContext(ctx context.Context) string {
	v, _ := ctx.Value(RequestIDKey).(string)
	return v
}

func NewUnaryUUID() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		id, err := newUUID()
		if err != nil {
			grpclog.Warningln(err.Error())
		}
		ctx = withUUID(ctx, id)
		return handler(ctx, req)

	}
}

func NewStreamUUID() grpc.StreamServerInterceptor {
	id, err := newUUID()
	if err != nil {
		grpclog.Warningln(err.Error())
	}
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		ctx := ss.Context()
		ctx = withUUID(ctx, id)
		stream := util.NewServerStreamWithContext(ss, ctx)
		return handler(srv, stream)
	}
}

// With returns a new context containing the supplied request ID value.  ctx is used as the parent of the resulting context.
func withUUID(ctx context.Context, xrid string) context.Context {
	if ctx == nil {
		ctx = context.TODO()
	}
	return context.WithValue(ctx, RequestIDKey, xrid)
}
