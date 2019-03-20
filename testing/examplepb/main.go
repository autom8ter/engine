package main

import (
	"github.com/autom8ter/engine"
	"github.com/autom8ter/engine/config"
	"github.com/autom8ter/engine/plugin"
	"log"
)

var examplePath = "example.plugin"

func main() {
	if err := engine.New().With(config.WithPluginLoaders(plugin.PluginLoader{
		Path:   examplePath,
		Symbol: "ExamplePb",
	})).Serve(); err != nil {
		log.Fatal(err.Error())
	}
}

