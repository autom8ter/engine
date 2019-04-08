package main

import (
	"github.com/autom8ter/engine"
	"github.com/autom8ter/engine/config"
	"github.com/autom8ter/engine/examples/examplepb"
	"log"
)

func main() {
	if err := engine.New("tcp", ":8080", true).With(
		//general options:
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
		config.WithStreamTraceMiddleware(),    // adds a streaming opentracing middleware
		config.WithStreamLoggingMiddleware(),  //adds a streaming logging middleware
		config.WithStreamRecoveryMiddleware(), // adds a streaming recovery middleware

	).Serve(); err != nil {
		log.Fatalln(err.Error())
	}
}

/*

 */
