package util

import (
	"google.golang.org/grpc/grpclog"
	"os"
	"strings"
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