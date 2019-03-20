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
	"fmt"
	"github.com/autom8ter/util"
	"github.com/spf13/cobra"
	"log"
	"os"
	"os/exec"
)

var args string
var config bool

type Output struct {
	name   string
	args   []string
	dir    string
	output string
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use: "protoctl",
	Long: `

  
.#####...#####....####...######...####....####...######..##.....
.##..##..##..##..##..##....##....##..##..##..##....##....##.....
.#####...#####...##..##....##....##..##..##........##....##.....
.##......##..##..##..##....##....##..##..##..##....##....##.....
.##......##..##...####.....##.....####....####.....##....######.
................................................................



`,
	Run: func(cmd *cobra.Command, _ []string) {
		if config {
			fmt.Println(configStr)
		} else {
			c := exec.Command("bash", "-c", "docker run -v `pwd`:/defs colemanword/prototool "+args)
			o := &Output{
				name: "docker",
				args: c.Args,
				dir:  c.Dir,
			}
			bit, err := c.CombinedOutput()
			if err != nil {
				log.Fatal(err.Error())
			}
			o.output = fmt.Sprintf("%s", bit)
			fmt.Println(util.ToPrettyJsonString(o))
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().StringVarP(&args, "args", "a", "generate", "arguments to pass to prototool")
	rootCmd.Flags().BoolVarP(&config, "config", "c", false, "print a prototool config to StdOut")
}

var configStr = `
protoc:
  version: 3.6.1
  allow_unused_imports: true
  includes:
    - /usr/local/include

lint:
  rules:
    remove:
      - FILE_OPTIONS_REQUIRE_JAVA_MULTIPLE_FILES
      - FILE_OPTIONS_REQUIRE_JAVA_PACKAGE
      - FILE_OPTIONS_REQUIRE_JAVA_OUTER_CLASSNAME
      - FILE_OPTIONS_EQUAL_GO_PACKAGE_PB_SUFFIX

generate:
  go_options:
    import_path: github.com/autom8ter/engine/testing

  plugins:
    - name: go
      type: go
      flags: plugins=grpc
      output: .

    - name: grpc-gateway
      type: go
      flags: logtostderr=true
      output: .

    - name: swagger
      type: go
      output: .
      flags: logtostderr=true

    - flags: binary,import_style=commonjs
      name: js
      output: client

    - name: cobra
      type: go
      output: .
      flags: plugins=client
`
