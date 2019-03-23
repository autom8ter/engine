# Engine

A Pluggable gRPC Microservice Framework
               
`go get github.com/autom8ter/engine`

`go get github.com/autom8ter/engine/enginectl`

Contributers: Coleman Word

License: MIT

```go
//Exported variable named Plugin, used to build with go/plugin
var Plugin  Example

//Embeded driver.PluginFunc is used to satisfy the driver.Plugin interface
type Example struct {
	driver.PluginFunc
}
//
func NewExample() Example {
	e := Example{}
	e.PluginFunc = func(s *grpc.Server) {
		examplepb.RegisterEchoServiceServer(s, e)
	}
	return e
}
//examplepb methods excluded for brevity

//A basic example with all config options:
func main() {
	if err := engine.New().With(
		//Look for configuration vars in environment with a set prefix.
	    config.WithEnvPrefix("ENGINE"),
		//hard coded plugins(not using go/plugins
		config.WithGoPlugins(),
		//tcp/unix and port/file, Only necessary if not using a config file(./config.json|config.yaml),  defaults to tcp, :3000
		config.WithGRPCListener("tcp", ":8080"),
		//Only necessary if not using a config file(./config.json|config.yaml) (variadic) no default
		//how to build: go build -buildmode=plugin -o ./plugins/$TARGETOUTPUT.plugin $TARGETGOFILE.go ref: https://golang.org/pkg/plugin/
		config.WithPluginPaths("bin/example.plugin"),
		//Only necessary if not using a config file(./config.json|config.yaml),  defaults to Plugin
		config.WithPluginSymbol("Plugin"),
		//add grpc server options (variadic) 
		config.WithServerOptions(),
		//add stream interceptors to all plugins(variadic) metrics, tracing, retry, auth, etc
		config.WithStreamInterceptors(),
		//add unary interceptors to all plugins(variadic) metrics, tracing, retry, auth, etc
		config.WithUnaryInterceptors(),
		//Only necessary if not using a config file(./config.json|config.yaml)
		config.WithDebug(),

	).Serve(); err != nil {
		//start server and fail if error
		grpclog.Fatal(err.Error())
	}
}
```
## Table of Contents

- [Engine](#engine)
  * [Overview](#overview)
  * [Features/Scope/Roadmap](#features-scope-roadmap)
    + [Engine Library:](#engine-library-)
    + [Enginectl (cli)](#enginectl--cli-)
  * [Plugin Interface](#plugin-interface)
  * [Configuration (viper)](#configuration--viper-)
  * [Grpc Middlewares (Guide)](#grpc-middlewares--guide-)
    + [Key Functions:](#key-functions-)
    + [Example(recovery):](#example-recovery--)
  * [EngineCtl (cli)](#enginectl--cli-)
  * [Limitations](#limitations)
  

## Overview

- Engine serves [go/plugins](https://golang.org/pkg/plugin/) that are dynamically loaded at runtime.
- Plugins export a type that implements the driver.Plugin interface: RegisterWithServer(s *grpc.Server)
- Exported plugins must be named Plugin
- Engine decouples the server runtime from grpc service development so that plugins can be added as externally compiled files that can be added to a deployment from local storage without making changes to the engine runtime.

## Features/Scope/Roadmap

### Engine Library:
- [x] Load grpc services from go/plugins at runtime that satisfy driver.Plugin
- [x] Load go/plugins from paths set in config file
- [x] Support for loading driver.Plugins directly(no go/plugins)
- [x] Support for custom gRPC Server options
- [x] Support for custom and chained Unary Interceptors
- [x] Support for custom and chained Stream Interceptors
- [ ] GoDoc documentation for every exported Method
- [ ] 80%+ code coverage
- [ ] Load go/plugins from paths set in environmental variables
- [ ] Load go/plugins directly from AWS S3
- [ ] Load go/plugins directly from GCP storage
- [ ] Load go/plugins directly from Github
- [ ] Load go/plugins directly from a Kubernetes Volume

### Enginectl (cli)
`go get github.com/autom8ter/engine/enginectl`
- [x] Load and serve grpc services from go/plugins at runtime that satisfy driver.Plugin
- [ ] Codegen: Makefile
- [ ] Codegen: Basic config file
- [ ] Codegen: Basic Protobuf file
- [ ] Codegen: Helm Chart
- [ ] Codegen: Dockerfile
- [ ] Codegen: Kubernetes Deployment
- [ ] Codegen: Google Endpoints Deployment
- [ ] Codegen: AWS API Gateway Deployment

## Driver

driver.Plugin is used to register grpc server implementations.

```go

//Plugin is an interface for representing gRPC server implementations.
type Plugin interface {
	RegisterWithServer(*grpc.Server)
}

//PluginFunc implements the Plugin interface.
type PluginFunc func(*grpc.Server)

//RegisterWithServer is an interface for representing gRPC server implementations.
func (p PluginFunc) RegisterWithServer(s *grpc.Server) {
	p(s)
}

```

## Configuration (viper)

- Config files must be either config.json or config.yaml in your current working directory

Variables:
- address: address to listen on, default: :3000
- network: tcp/unix, default: tcp
- paths: paths to generated plugin files to load
- symbol: exported plugin variable name, default: Plugin
- debug: enable verbose logging for development

example:
```json
{
  "address": ":3000",
  "network": "tcp",
  "paths": [
    "bin/example.plugin"
  ],
  "symbol": "Plugin",
  "debug": "true"
}

```

## Grpc Middlewares (Guide)

### Key Functions:
    type StreamServerInterceptor func(srv interface{}, ss ServerStream, info *StreamServerInfo, handler StreamHandler) error
    type UnaryServerInterceptor func(ctx context.Context, req interface{}, info *UnaryServerInfo, handler UnaryHandler) (resp interface{}, err error)

### Example(recovery): 
```go
// UnaryServerInterceptor returns a new unary server interceptor for panic recovery.
func UnaryServerInterceptor(opts ...Option) grpc.UnaryServerInterceptor {
	o := evaluateOptions(opts)
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (_ interface{}, err error) {
		defer func() {
			if r := recover(); r != nil {
				err = recoverFrom(ctx, r, o.recoveryHandlerFunc)
			}
		}()

		return handler(ctx, req)
	}
}

// StreamServerInterceptor returns a new streaming server interceptor for panic recovery.
func StreamServerInterceptor(opts ...Option) grpc.StreamServerInterceptor {
	o := evaluateOptions(opts)
	return func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
		defer func() {
			if r := recover(); r != nil {
				err = recoverFrom(stream.Context(), r, o.recoveryHandlerFunc)
			}
		}()

		return handler(srv, stream)
	}
}
```
Please see [go-grpc-middleware](https://github.com/grpc-ecosystem/go-grpc-middleware) for a list of useful
Unary and Streaming Interceptors

## EngineCtl (cli)

EngineCtl is a very basic implementation of the Engine library that allows
users to use flags to override a config file. It requires zero coding to 
use most of the functionality of the engine library since users only need to 
provide paths to plugins to create a fully customizable grpc microservice.

**It is particularly useful in containers:**

-> download enginectl -> copy plugins to container->copy config file to container-> enginectl init

Run `enginectl` with no flags/subcommands

Output:

```text

----------------------------------------------------------------------------
8888888888                d8b                        888   888
888                       Y8P                        888   888
888                                                  888   888
8888888   88888b.  .d88b. 88888888b.  .d88b.  .d8888b888888888
888       888 "88bd88P"88b888888 "88bd8P  Y8bd88P"   888   888
888       888  888888  888888888  88888888888888     888   888
888       888  888Y88b 888888888  888Y8b.    Y88b.   Y88b. 888
8888888888888  888 "Y88888888888  888 "Y8888  "Y8888P "Y888888
                       888                                    
                  Y8b d88P                                    
                   "Y88P"

Assign individual developers to develop specific plugins and then 
just add them as a plugin config path. Plugin development is completely
independent of the runtime NICE.

----------------------------------------------------------------------------
Download:
go get github.com/autom8ter/engine/enginectl

----------------------------------------------------------------------------
Configuration:

- Config files must be either config.json or config.yaml in your current working directory

Variables:
- address: address to listen on, default: :3000
- network: tcp/unix, default: tcp
- paths: paths to generated plugin files to load
- symbol: exported plugin variable name, default: Plugin
- debug: enable verbose logging for development

example:

{
  "address": ":3000",
  "network": "tcp",
  "paths": [
    "bin/example.plugin"
  ],
  "symbol": "Plugin",
  "debug": "true"
}

----------------------------------------------------------------------------
How to build go/plugins:
go build -buildmode=plugin -o ../bin/example.plugin examplepb/plugin.go
----------------------------------------------------------------------------
Example Dockerfile

FROM golang:1.11
RUN go get github.com/autom8ter/engine/enginectl
COPY bin/example.plugin .
COPY config.json .
ENTRYPOINT [ "enginectl", "init"] 

----------------------------------------------------------------------------

Current Config:
map[paths:[bin/example.plugin] symbol:Plugin debug:true address::3000 network:tcp]
----------------------------------------------------------------------------

Usage:
  enginectl [command]

Available Commands:
  help        Help about any command
  serve       load plugins from config and start the enginectl server

Flags:
  -h, --help   help for enginectl

Use "enginectl [command] --help" for more information about a command.



```
    output:
    INFO: 2019/03/22 16:24:14 using config file: /Users/coleman/go/src/github.com/autom8ter/engine/config.json
    INFO: 2019/03/22 16:24:14 creating server config from config file
    INFO: 2019/03/22 16:24:14 registered paths: [bin/example.plugin]
    INFO: 2019/03/22 16:24:14 registered plugin: *main.Example
    INFO: 2019/03/22 16:24:14 loading environmental variables with prefix: 
    INFO: 2019/03/22 16:24:14 setting environmental key replacer: replace: . with: _
    INFO: 2019/03/22 16:24:14 set plugin symbol: Plugin
    INFO: 2019/03/22 16:24:14 creating grpc server
    INFO: 2019/03/22 16:24:14 registered server reflection
    INFO: 2019/03/22 16:24:14 plugin count: 1
    INFO: 2019/03/22 16:24:14 creating server listener tcp :3000
    INFO: 2019/03/22 16:24:14 gRPC server is starting [::]:3000

## Limitations

Im hoping someone can help explain why some of these errors occur:
- When creating a plugin, one must NOT use pointer methods when satisfying the driver.Plugin interface
- If a json config is hard-coded as a string, the server fails, but succeeds if it is present as a config file