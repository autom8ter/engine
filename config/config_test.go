package config_test

import (
	"errors"
	"github.com/autom8ter/engine/config"
	"os"
	"testing"
)

func TestNew(t *testing.T) {
	c := config.New("tcp", ":4000")
	if c == nil {
		t.Fatal(errors.New("nil config"))
	}
	if c.Network != "tcp" {
		t.Fatal(errors.New("expected tcp"))
	}
	if c.Address != ":4000" {
		t.Fatal(errors.New("expected :4000"))
	}
	if c.Symbol != "Plugin" {
		t.Fatal(errors.New("expected Plugin"))
	}
	c.With(
		config.WithDebug(),
	)
	if os.Getenv("DEBUG") != "t" {
		t.Fatal(errors.New("expected t"))
	}
}
