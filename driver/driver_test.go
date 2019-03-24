package driver_test

import (
	"github.com/autom8ter/engine/driver"
	"github.com/autom8ter/engine/examples/examplepb/mock"
	"google.golang.org/grpc"
	"testing"
)

var ex = mock.NewExample()

func TestIsPlugin(t *testing.T) {
	if !driver.IsPlugin(ex) {
		t.Fatal("not a plugin")
	}
}

func TestNewPlugin(t *testing.T) {
	this := driver.NewPlugin(ex.GRPCFunc)
	if !driver.IsPlugin(this) {
		t.Fatal("not a plugin")
	}
}

func TestNewPluginFunc(t *testing.T) {
	this := driver.NewGRPCFunc(ex.GRPCFunc)
	if !driver.IsPlugin(this) {
		t.Fatal("not a plugin")
	}
}

func TestPluginFunc_RegisterWithServer(t *testing.T) {
	ex.GRPCFunc.RegisterWithServer(grpc.NewServer())
}
