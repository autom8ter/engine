package main

import (
	"github.com/autom8ter/engine"
	"github.com/autom8ter/engine/config"
	"log"
)

func main() {
	e := engine.New("tcp", ":8080").With(
		config.WithDebug(), //adds verbose logging for development
		config.WithPlugins("Plugin", "bin/example.so"), //loads a plugin with the exported symbol from the specified path
		config.WithChannelz(), //adds a channelz service
		config.WithReflection(),//adds a reflection service
		config.WithHealthz(),//adds a healthz service
		config.WithUnaryLogger(),//adds a unary logger
		config.WithStreamLogger(),//adss a streaming logger
		config.WithMaxConcurrentStreams(1000),//sets max concurrent server streams
	)
	defer e.Shutdown()
	if err := e.Serve(); err != nil {
		log.Fatalln(err.Error())
	}
}
