package engine_test

import (
	"github.com/autom8ter/engine"
	"github.com/autom8ter/engine/config"
	"github.com/autom8ter/engine/plugin"
	"os"
	"testing"
)
var home = os.Getenv("HOME")

var examplePath = home+"/.plugins/example.plugin"

func TestNewEngine(t *testing.T) {
	if err := engine.New().With(config.WithPluginLoaders(plugin.PluginLoader{
		Path:   examplePath,
		Symbol: "ExamplePb",
	})).Serve(); err != nil {
		t.Fatal(err.Error())
	}
}
