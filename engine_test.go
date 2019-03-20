package engine_test

import (
	"github.com/autom8ter/engine"
	"github.com/autom8ter/engine/config"
	"testing"
)

func TestNewEngine(t *testing.T) {
	if err := engine.New().With(
		config.WithDefaultLogger(),
		config.WithAddr("tcp", ":8080"),
		config.WithPlugins(),
	).Serve(); err != nil {
		t.Fatal(err.Error())
	}
}
