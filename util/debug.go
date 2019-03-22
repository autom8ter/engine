package util

import (
	"github.com/spf13/viper"
	"google.golang.org/grpc/grpclog"
)

func Debugf(format string, args ...interface{}) {
	if viper.GetBool("debug") {
		grpclog.Infof(format, args...)
	}
}

func Debugln(args ...interface{}) {
	if viper.GetBool("debug") {
		grpclog.Infoln(args...)
	}
}
