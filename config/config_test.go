package config_test

import (
	"github.com/autom8ter/engine/config"
	"github.com/spf13/viper"
	"reflect"
	"testing"
)

var c = config.New()

func TestNew(t *testing.T) {
	expect("plugin", "Plugin", c.Symbol, t)
	expect("network", "tcp", c.Network, t)
	expect("address", ":3000", c.Address, t)
}

func TestWith(t *testing.T) {
	TestNew(t)
	c.With(
		config.WithDebug(),
		config.WithEnvPrefix("ENGINE"),
		config.WithPluginSymbol("Random"),
		config.WithNetwork("tcp", ":3001"),
	)

	expect("env_prefix", "ENGINE", viper.GetString("env_prefix"), t)
	expect("debug", true, viper.GetBool("debug"), t)
	expect("symbol", "Random", viper.GetString("symbol"), t)
	expect("network", "tcp", viper.GetString("network"), t)
	expect("address", ":3001", viper.GetString("address"), t)

}

func expect(key string, exp, got interface{}, t *testing.T) {
	if !reflect.DeepEqual(exp, got) {
		t.Fatalf("key:%s expected: %s got: %s\n", key, exp, got)
	}
}
