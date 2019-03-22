package util

import (
	"github.com/spf13/viper"
	"google.golang.org/grpc/grpclog"
)

// Debugf is grpclog.Infof(format, args...) but only executes if debug=true is set in your config or environmental variables
func Debugf(format string, args ...interface{}) {
	if viper.GetBool("debug") {
		grpclog.Infof(format, args...)
	}
}

// Debugln is grpclog.Infoln(args...) but only executes if debug=true is set in your config or environmental variables
func Debugln(args ...interface{}) {
	if viper.GetBool("debug") {
		grpclog.Infoln(args...)
	}
}
