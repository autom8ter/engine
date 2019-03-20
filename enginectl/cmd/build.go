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
	"github.com/autom8ter/engine/plugin"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"google.golang.org/grpc/grpclog"
)

var file string

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "build and store a plugin",
	Run: func(cmd *cobra.Command, args []string) {
		if file == "" {
			grpclog.Fatalln(errors.New("please use the --file || -f flag to specify a file to build a plugin for"))
		}

		bits, err := plugin.Build(file)
		if err != nil {
			grpclog.Fatalln(err.Error())
		}
		grpclog.Infof("Build Plugin Output: \n%s\n", bits)
	},
}

func init() {
	buildCmd.Flags().StringVarP(&file, "file", "f", "", "file to generate plugin for")
	rootCmd.AddCommand(buildCmd)
}
