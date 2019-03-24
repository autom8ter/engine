package main

import (
	"github.com/autom8ter/engine"
	"github.com/autom8ter/engine/config"
	"log"
)

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
		config.WithUnaryLoggingMiddleware(),  // adds a unary logging rmiddleware
		config.WithUnaryRecoveryMiddleware(), // adds a unary recovery middleware
		config.WithUnaryTraceMiddleware(),    // adds a streaming opentracing middleware

		//streaming middleware
		config.WithStreamLoggingMiddleware(),  //adds a streaming logging middleware
		config.WithStreamRecoveryMiddleware(), // adds a streaming recovery middleware
		config.WithStreamTraceMiddleware(),    // adds a streaming opentracing middleware

	).Serve(); err != nil {
		log.Fatalln(err.Error())
	}
}
