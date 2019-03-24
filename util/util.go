package util

import (
	"context"
	"google.golang.org/grpc/grpclog"
	"os"
	"os/signal"
	"plugin"
	"strings"
	"syscall"
)

// Debugf is grpclog.Infof(format, args...) but only executes if debug=true is set in your config or environmental variables
func Debugf(format string, args ...interface{}) {
	if strings.Contains(os.Getenv("DEBUG"), "t") || strings.Contains(os.Getenv("DEBUG"), "T") {
		grpclog.Infof(format, args...)
	}
}

// Debugln is grpclog.Infoln(args...) but only executes if debug=true is set in your config or environmental variables
func Debugln(args ...interface{}) {
	if strings.Contains(os.Getenv("DEBUG"), "t") || strings.Contains(os.Getenv("DEBUG"), "T") {
		grpclog.Infoln(args...)
	}
}

func Load(path, symbol string) interface{} {
	plug, err := plugin.Open(path)
	if err != nil {
		grpclog.Fatalln(err.Error())
	}
	sym, err := plug.Lookup(symbol)
	if err != nil {
		grpclog.Fatalln(err.Error())
	}
	return sym
}

func CancelFunc(ctx context.Context, cancel func()) func() error {
	return func() error {
		sdCh := make(chan os.Signal, 1)
		defer close(sdCh)
		defer signal.Stop(sdCh)
		signal.Notify(sdCh, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
		select {
		case <-sdCh:
			cancel()
		case <-ctx.Done():
			// no-op
		}
		return nil
	}
}
