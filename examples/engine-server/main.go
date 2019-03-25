package main

import (
	"context"
	"github.com/autom8ter/engine"
	"github.com/autom8ter/engine/config"
	"github.com/autom8ter/engine/lib"
	"log"
)

func main() {
	dbctx := context.Background()
	mongoClient := lib.MongoClient("mongodb://localhost:27017", dbctx)

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
		config.WithUnaryUUIDMiddleware(),                                                     //adds a unary uuid middleware
		config.WithUnaryPingMongoMiddleware(mongoClient, dbctx),                              //ping mongo db
		config.WithUnarySaveToMongoMiddleware(mongoClient, "testdb", "testcoll", "document"), //adds the request object to mongodb under the given dbName, collection, and key extracted from the request context
		config.WithUnaryTraceMiddleware(),                                                    // adds a streaming opentracing middleware
		config.WithUnaryLoggingMiddleware(),                                                  // adds a unary logging rmiddleware
		config.WithUnaryRecoveryMiddleware(),                                                 // adds a unary recovery middleware

		//streaming middleware
		config.WithStreamUUIDMiddleware(),     //adds a streaming uuid middleware
		config.WithStreamTraceMiddleware(),    // adds a streaming opentracing middleware
		config.WithStreamLoggingMiddleware(),  //adds a streaming logging middleware
		config.WithStreamRecoveryMiddleware(), // adds a streaming recovery middleware

	).Serve(); err != nil {
		log.Fatalln(err.Error())
	}
}
