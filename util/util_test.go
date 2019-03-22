package util_test

import (
	"github.com/autom8ter/engine/util"
	"github.com/spf13/viper"
	"testing"
)

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
