# Engine

A Pluggable gRPC Microservice Framework
               
`go get github.com/autom8ter/engine`

Contributers: Coleman Word

License: MIT

```go
//Exported variable named Plugin, used to build with go/plugin
//Compile plugin and add to your config path to be loaded by the engine instance
// ex: go build -buildmode=plugin -o bin/example.so examplepb/plugin.go
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

//The compiled plugin file will be loaded at runtime if its set in your config path.
//example:
func main() {
	if err := engine.New("tcp", ":8080").With(
		//general options:
		config.WithDebug(),                    //adds verbose logging for development
		config.WithMaxConcurrentStreams(1000), //sets max concurrent server streams

		//plugins:
		config.WithPlugins("Plugin", "bin/example.so"), //loads a plugin with the exported symbol from the specified path
		config.WithChannelz(),                          //adds a channelz service
		config.WithReflection(),                        //adds a reflection service
		config.WithHealthz(),                           //adds a healthz service

		//unary middleware:
		config.WithUnaryUUIDMiddleware(),     //adds a unary uuid middleware
		config.WithUnaryTraceMiddleware(),    // adds a streaming opentracing middleware
		config.WithUnaryLoggingMiddleware(),  // adds a unary logging rmiddleware
		config.WithUnaryRecoveryMiddleware(), // adds a unary recovery middleware

		//streaming middleware
		config.WithStreamUUIDMiddleware(),     //adds a streaming uuid middleware
		config.WithStreamTraceMiddleware(),    // adds a streaming opentracing middleware
		config.WithStreamLoggingMiddleware(),  //adds a streaming logging middleware
		config.WithStreamRecoveryMiddleware(), // adds a streaming recovery middleware

	).Serve(); err != nil {
		log.Fatalln(err.Error())
	}
}

/*
Output:
------------------------------------------------
         #                    #               
         ##                   ##              
######## ###  ##   ###### ### ###  ## ########
         #### ##  ###     ### #### ##         
 ####### #######  ###  ## ### #######  #######
 ###     ### ###  ###  ## ### ### ###  ###    
 ####### ###  ##   ###### ### ###  ##  #######
               #                    #
Unary_Interceptors: 4
Stream_Interceptors: 4
Server_Options: 1
Plugins: 4
Plugin_Paths: [bin/example.so]
Plugin_Symbol: Plugin
Network: tcp
Address: :8080
------------------------------------------------


*/

```
---

## Table of Contents

- [Engine](#engine)
  * [Table of Contents](#table-of-contents)
  * [Overview](#overview)
  * [Features/Scope/Roadmap](#features-scope-roadmap)
  * [Driver](#driver)
  * [Grpc Middlewares](#grpc-middlewares)
    + [Key Functions:](#key-functions-)
    + [Example(recovery):](#example-recovery--)
  * [GoDoc](#godoc)
      - [type Engine](#type-engine)
      - [func  New](#func--new)
      - [type Runtime](#type-runtime)
      - [func (*Runtime) Config](#func---runtime--config)
      - [func (*Runtime) Serve](#func---runtime--serve)
      - [func (*Runtime) Shutdown](#func---runtime--shutdown)
      - [func (*Runtime) With](#func---runtime--with)
  * [Limitations](#limitations)
  
---

## Overview

- Engine serves [go/plugins](https://golang.org/pkg/plugin/) that are dynamically loaded at runtime.
- Plugins must export a type that implements the driver.Plugin interface: RegisterWithServer(s *grpc.Server)
- Engine decouples the server runtime from grpc service development so that plugins can be added as externally compiled files that can be added to a deployment from local storage without making changes to the engine server.

---

## Features/Scope/Roadmap

- [x] Load grpc services from go/plugins at runtime that satisfy driver.Plugin
- [x] Support for loading driver.Plugins directly(no go/plugins)
- [x] Support for custom gRPC Server options
- [x] Support for custom and chained Unary Interceptors
- [x] Support for custom and chained Stream Interceptors
- [x] GoDoc documentation for every exported Method
- [x] Channelz service option ref: https://godoc.org/google.golang.org/grpc/channelz
- [x] Reflection service option ref: https://godoc.org/google.golang.org/grpc/reflection
- [x] Healthz service option ref: https://godoc.org/google.golang.org/grpc/health
- [x] Unary logger middleware option
- [x] Unary recovery middleware option
- [x] Unary tracing middleware option
- [x] Stream logger middleware option
- [x] Stream recovery middleware option
- [x] Stream tracing middleware option
- [x] Unary metrics middleware option
- [x] Stream metrics middleware option


---

## Driver

github.com/autom8ter/engine/driver

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
---

## Limitations

Im hoping someone can help explain why some of these errors occur:
- When creating a plugin, one must NOT use pointer methods when satisfying the driver.Plugin interface
- If a json config is hard-coded as a string, the server fails, but succeeds if it is present as a config file

---

