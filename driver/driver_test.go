package driver_test

import (
	"github.com/autom8ter/engine/driver"
	"github.com/autom8ter/engine/examples/examplepb"
	"google.golang.org/grpc"
	"testing"
)

var ex = examplepb.NewExample()

func TestIsPlugin(t *testing.T) {
	if !driver.IsPlugin(ex) {
		t.Fatal("not a plugin")
	}
}

func TestNewPlugin(t *testing.T) {
	this := driver.NewPlugin(ex.PluginFunc)
	if !driver.IsPlugin(this) {
		t.Fatal("not a plugin")
	}
}

func TestNewPluginFunc(t *testing.T) {
	this := driver.NewPluginFunc(ex.PluginFunc)
	if !driver.IsPlugin(this) {
		t.Fatal("not a plugin")
	}
}

func TestPluginFunc_RegisterWithServer(t *testing.T) {
	ex.PluginFunc.RegisterWithServer(grpc.NewServer())
}
