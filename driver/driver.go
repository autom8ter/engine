package driver

import (
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"os"
)

func init() {
	grpclog.SetLoggerV2(grpclog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout))
	viper.AutomaticEnv()
	viper.AddConfigPath(".")
	viper.SetDefault("network", "tcp")
	viper.SetDefault("address", ":3000")
	if err := viper.ReadInConfig(); err != nil {
		grpclog.Warningln(err.Error())
	} else {
		grpclog.Infof("using config file: %s\n", viper.ConfigFileUsed())
	}
}

//Plugin is an interface for representing gRPC server implementations.
type Plugin interface {
	RegisterWithServer(*grpc.Server)
}

type PluginFunc func(s *grpc.Server)

func (p PluginFunc) RegisterWithServer(s *grpc.Server) {
	p(s)
}
