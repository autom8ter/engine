package lib

import (
	"context"
	"github.com/juju/ratelimit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

const infinityDuration time.Duration = 0x7fffffffffffffff

type RateLimitOption func(*rateLimiter)

// WithLimiter customizes your limiter in the middleware

func WithLimiter(l Limiter) RateLimitOption {

	return func(r *rateLimiter) {

		r.limiter = l

		if r.maxWaitDuration == 0 {

			r.maxWaitDuration = infinityDuration

		}

	}

}

// WithMaxWaitDuration customizes maxWaitDuration in limiter's WaitMaxDuration action.

func WithMaxWaitDuration(maxWaitDuration time.Duration) RateLimitOption {

	return func(r *rateLimiter) {

		r.maxWaitDuration = maxWaitDuration

	}

}

type tokenBucketLimiter struct {
	limiter *ratelimit.Bucket
}

// NewTokenBucketRateLimiter creates a tokenBucketLimiter.

func NewTokenBucketRateLimiter(fillInterval time.Duration, capacity, quantum int64) *tokenBucketLimiter {

	return &tokenBucketLimiter{

		limiter: ratelimit.NewBucketWithQuantum(fillInterval, capacity, quantum),
	}

}

// WaitMaxDuration

func (b *tokenBucketLimiter) WaitMaxDuration(maxWaitDuration time.Duration) bool {

	return b.limiter.WaitMaxDuration(1, maxWaitDuration)
}

// StreamServerInterceptor returns a new stream server interceptors that performs request rate limit.

func NewStreamRateLimiter(opts ...RateLimitOption) grpc.StreamServerInterceptor {

	ratelimiter := emptyRatelimiter()

	for _, opt := range opts {

		opt(ratelimiter)

	}

	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {

		if ratelimiter.Wait() {

			return handler(srv, stream)

		}

		return status.Errorf(codes.ResourceExhausted, "%s is rejected by grpc_ratelimit middleare, please retry later.", info.FullMethod)

	}

}

type Limiter interface {
	WaitMaxDuration(time.Duration) bool
}

type rateLimiter struct {
	limiter Limiter

	maxWaitDuration time.Duration
}

func (r *rateLimiter) Wait() bool {

	return r.limiter.WaitMaxDuration(r.maxWaitDuration)

}

type emptyLimiter struct{}

func (e *emptyLimiter) WaitMaxDuration(time.Duration) bool {

	return true

}

func emptyRatelimiter() *rateLimiter {

	return &rateLimiter{

		limiter: &emptyLimiter{},
	}

}

// NewUnaryRateLimiter returns a new unary server interceptors that performs request rate limit.

func NewUnaryRateLimiter(opts ...RateLimitOption) grpc.UnaryServerInterceptor {
	ratelimiter := emptyRatelimiter()

	for _, opt := range opts {

		opt(ratelimiter)

	}
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {

		if ratelimiter.Wait() {

			return handler(ctx, req)

		}

		return nil, status.Errorf(codes.ResourceExhausted, "%s is rejected by grpc_ratelimit middleare, please retry later.", info.FullMethod)

	}
}
