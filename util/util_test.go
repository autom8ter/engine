package util_test

import (
	"github.com/autom8ter/engine/util"
	"github.com/spf13/viper"
	"testing"
)

func init() {
	if !viper.InConfig("symbol") {
		viper.SetDefault("symbol", "Plugin")
	}
	if !viper.InConfig("address") {
		viper.SetDefault("address", ":3000")
	}
}

func TestChannelzClient(t *testing.T) {
	cc := util.ChannelzClient(viper.GetString("address"))
	if cc == nil {
		t.Fatal("nil channelz client")
	}
}

func TestDebugf(t *testing.T) {
	viper.Set("debug", true)
	if !viper.GetBool("debug") {
		t.Fatal("expected debug true")
	}
	viper.Set("debug", false)
	if viper.GetBool("debug") {
		t.Fatal("expected debug false")
	}
}

func TestDebugln(t *testing.T) {
	viper.Set("debug", true)
	if !viper.GetBool("debug") {
		t.Fatal("expected debug true")
	}
	viper.Set("debug", false)
	if viper.GetBool("debug") {
		t.Fatal("expected debug false")
	}
}

func TestLoadPlugins(t *testing.T) {
	viper.Set("paths", []string{"../bin/example.plugin"})
	plugs := util.LoadPlugins()
	if len(plugs) == 0 {
		t.Fatal("expected at least one plugin")
	}
}
