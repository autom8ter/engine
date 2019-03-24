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
		config.WithUnaryUUIDMiddleware(),  //adds a unary uuid middleware
		config.WithUnaryTraceMiddleware(),    // adds a streaming opentracing middleware
		config.WithUnaryLoggingMiddleware(),  // adds a unary logging rmiddleware
		config.WithUnaryRecoveryMiddleware(), // adds a unary recovery middleware

		//streaming middleware
		config.WithStreamUUIDMiddleware(), //adds a streaming uuid middleware
		config.WithStreamTraceMiddleware(),    // adds a streaming opentracing middleware
		config.WithStreamLoggingMiddleware(),  //adds a streaming logging middleware
		config.WithStreamRecoveryMiddleware(), // adds a streaming recovery middleware

	).Serve(); err != nil {
		log.Fatalln(err.Error())
	}
}

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

- [ ] Unary metrics middleware option
- [ ] Stream metrics middleware option

- [ ] Auth middleware option

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

## Grpc Middlewares

Middlewares should be used for things like monitoring, logging, auth, retry, etc.

They can be added to the engine with:

    config.WithStreamInterceptors(...)
    config.WithUnaryInterceptors(...)

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

---

## GoDoc

#### type Engine

```go
type Engine interface {
	With(opts ...config.Option) *Runtime
	Config() *config.Config
	Shutdown()
	Serve() error
}
```
Engine is an interface used to describe a server runtime


#### func  New

```go
func New(network, addr, symbol string, paths ...string) Engine
```
New creates a engine intstance.

#### type Runtime

```go
type Runtime struct {
}
```

Runtime is an implementation of the engine API.

#### func (*Runtime) Config

```go
func (e *Runtime) Config() *config.Config
```
Config returns the runtimes current configuration

#### func (*Runtime) Serve

```go
func (e *Runtime) Serve() error
```
Serve starts the runtime gRPC server.

#### func (*Runtime) Shutdown

```go
func (e *Runtime) Shutdown()
```
Shutdown gracefully closes the grpc server.

#### func (*Runtime) With

```go
func (e *Runtime) With(opts ...config.Option) *Runtime
```
With wraps the runtimes config with config options ref:
github.com/autom8ter/engine/config/options.go

---

## Limitations

Im hoping someone can help explain why some of these errors occur:
- When creating a plugin, one must NOT use pointer methods when satisfying the driver.Plugin interface
- If a json config is hard-coded as a string, the server fails, but succeeds if it is present as a config file

