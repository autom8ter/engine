package plugin

import (
	"fmt"
	"github.com/autom8ter/engine/driver"
	"github.com/spf13/viper"
	"google.golang.org/grpc/grpclog"
	"os"
	"plugin"
)

func LoadPlugins() []driver.Plugin {
	var plugs = []driver.Plugin{}
	for _, p := range viper.GetStringSlice("paths") {
		plug, err := plugin.Open(p)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		sym, err := plug.Lookup("Plugin")
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		var myPlugin driver.Plugin
		myPlugin, ok := sym.(driver.Plugin)
		if !ok {
			fmt.Printf("provided plugin: %T does not satisfy Plugin interface\n", sym)
			os.Exit(1)
		} else {
			grpclog.Infof("registered plugin: %T\n", sym)
			plugs = append(plugs, myPlugin)
		}
	}

	return plugs
}
