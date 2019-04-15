# Engine

[Godoc](https://github.com/autom8ter/engine/blob/master/GODOC.md)

A Pluggable gRPC Microservice Framework
               
`go get github.com/autom8ter/engine`

Contributers: Coleman Word

License: MIT

```go

//Embeded driver.PluginFunc is used to satisfy the driver.Plugin interface
type Example struct {
	driver.PluginFunc
}

//create new example with the generated registration function from grpc
func NewExample() Example {
	e := Example{}
	e.PluginFunc = func(s *grpc.Server) {
		examplepb.RegisterEchoServiceServer(s, e)
	}
	return e
}
//examplepb methods excluded for brevity

func main() {
	if err := engine.New("tcp", ":8080").With(
		//general options:
		config.WithDebug(),                    //adds verbose logging for development
		config.WithMaxConcurrentStreams(1000), //sets max concurrent server streams

		//plugins:
		config.WithChannelz(),   //adds a channelz service
		config.WithReflection(), //adds a reflection service
		config.WithHealthz(),    //adds a healthz service
		config.WithPlugins(examplepb.NewExample()),

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

```

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

## Features/Scope/Roadmap

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

