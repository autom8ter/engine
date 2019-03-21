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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc/grpclog"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "load plugins from $HOME/.plugins and start the enginectl server",
	Run: func(cmd *cobra.Command, args []string) {
		viper.Debug()
		if err := eng().Serve(); err != nil {
			grpclog.Fatalln(err.Error())
		}

	},
}

func eng() engine.Engine {
	return engine.New().With(
		config.WithGRPCLogger(),
	)
}
