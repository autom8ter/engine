// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"github.com/autom8ter/engine"
	"github.com/autom8ter/engine/config"
	"github.com/autom8ter/util"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc/grpclog"
)

var address string
var network string
var paths []string
var debug bool
var symbol string
var envPrefix string

func init() {
	var err error
	serveCmd.Flags().StringVarP(&address, "address", "a", viper.GetString("address"), "network address to listen on")
	serveCmd.Flags().StringVarP(&network, "network", "n", viper.GetString("network"), "network type to listen on")
	serveCmd.Flags().StringSliceVarP(&paths, "paths", "p", viper.GetStringSlice("paths"), "relative paths to plugins to register")
	serveCmd.Flags().BoolVarP(&debug, "debug", "d", false, "enable debug mode")
	serveCmd.Flags().StringVarP(&symbol, "symbol", "s", viper.GetString("symbol"), "plugin symbol to scan plugin path for")
	serveCmd.Flags().StringVarP(&envPrefix, "env_prefix", "e", viper.GetString("env_prefix"), "env prefix to set")

	if len(paths) == 0 || paths == nil || paths[0] == "" {
		paths, err = util.ReadAsCSV(util.Prompt("please provide path(s) to plugins: "))
		if err != nil {
			grpclog.Fatalf("please provide valid paths seperated by commas\n%s\n", err.Error())
		}
		if len(paths) == 0 || paths == nil || paths[0] == "" {
			grpclog.Fatalf("failed to register plugin paths, current config: \n%s\n", util.ToPrettyJsonString(viper.AllSettings()))
		}
	}
	if err := viper.BindPFlags(serveCmd.Flags()); err != nil {
		grpclog.Fatalln(err.Error())
	}
}

// initCmd represents the init command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "load plugins from config and start the enginectl server",
	Run: func(cmd *cobra.Command, args []string) {
		//A basic example with all config options:
		if err := engine.New().With(
			//tcp/unix and port/file, Only necessary if not using a config file(./config.json|config.yaml),  defaults to tcp, :3000
			config.WithNetwork(network, address),
			//Only necessary if not using a config file(./config.json|config.yaml) (variadic) no default
			config.WithPluginPaths(paths...),
			config.WithEnvPrefix(envPrefix),
			config.WithPluginSymbol(symbol),
			/*	config.WithServerOptions(),
				func ConnectionTimeout(d time.Duration) ServerOption
				func Creds(c credentials.TransportCredentials) ServerOption
				func CustomCodec(codec Codec) ServerOption
				func InTapHandle(h tap.ServerInHandle) ServerOption
				func InitialConnWindowSize(s int32) ServerOption
				func InitialWindowSize(s int32) ServerOption
				func KeepaliveEnforcementPolicy(kep keepalive.EnforcementPolicy) ServerOption
				func KeepaliveParams(kp keepalive.ServerParameters) ServerOption
				func MaxConcurrentStreams(n uint32) ServerOption
				func MaxHeaderListSize(s uint32) ServerOption
				func MaxMsgSize(m int) ServerOption
				func MaxRecvMsgSize(m int) ServerOption
				func MaxSendMsgSize(m int) ServerOption
				func RPCCompressor(cp Compressor) ServerOption
				func RPCDecompressor(dc Decompressor) ServerOption
				func ReadBufferSize(s int) ServerOption
				func StatsHandler(h stats.Handler) ServerOption
				func StreamInterceptor(i StreamServerInterceptor) ServerOption
				func UnaryInterceptor(i UnaryServerInterceptor) ServerOption
				func UnknownServiceHandler(streamHandler StreamHandler) ServerOption
				func WriteBufferSize(s int) ServerOption
			*/
		).Serve(); err != nil {
			//start server and fail if error
			grpclog.Fatal(err.Error())
		}
	},
}
