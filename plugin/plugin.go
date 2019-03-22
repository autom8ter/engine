package plugin

import (
	"github.com/autom8ter/engine/driver"
	"github.com/autom8ter/engine/lib/util"
	"github.com/spf13/viper"
	"google.golang.org/grpc/grpclog"
	"plugin"
)

func init() {
	viper.SetDefault("symbol", "Plugin")
}

func LoadPlugins() []driver.Plugin {
	var plugs = []driver.Plugin{}
	for _, p := range viper.GetStringSlice("paths") {

		util.Debugf("registered paths: %v\n", viper.GetStringSlice("paths"))
		plug, err := plugin.Open(p)
		if err != nil {
			grpclog.Fatalln(err.Error())
		}
		sym, err := plug.Lookup(viper.GetString("symbol"))
		if err != nil {
			grpclog.Fatalln(err.Error())
		}

		var asPlugin driver.Plugin
		asPlugin, ok := sym.(driver.Plugin)
		if !ok {
			grpclog.Fatalf("provided plugin: %T does not satisfy Plugin interface\n", sym)
		} else {
			util.Debugf("registered plugin: %T\n", sym)
			plugs = append(plugs, asPlugin)
		}
	}

	return plugs
}
