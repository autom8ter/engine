package engine_test

import (
	"github.com/autom8ter/engine"
	"github.com/autom8ter/engine/testing/helloworld"
	"testing"
)

func TestNewEngine(t *testing.T) {
	e := engine.New(
		engine.WithDefaultLogger(),
		engine.WithPlugins(helloworld.NewBasicGreeter()),
	)
	if err := e.Serve(); err != nil {
		t.Fatal(err.Error())
	}
}
