package engine_test

import (
	"github.com/autom8ter/engine"
	"github.com/autom8ter/engine/config"
	"github.com/autom8ter/engine/testing/examplepb"
	"testing"
)

func TestNewEngine(t *testing.T) {
	if err := engine.New(examplepb.NewExample()).With(config.WithDefaultLogger()).Serve(); err != nil {
		t.Fatal(err.Error())
	}
}
