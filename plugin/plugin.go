package plugin

import (
	"fmt"
	"github.com/autom8ter/engine/driver"
	"github.com/spf13/viper"
	"google.golang.org/grpc/grpclog"
	"plugin"
)

func init() {
	if err := viper.Unmarshal(loader); err != nil {
		grpclog.Fatalln(err.Error())
	}
}

type Loader struct {
	Plugins []string `mapstructure:"plugins" json:"plugins"`
}

var loader = &Loader{}

func GetLoader() *Loader {
	return loader
}

func (p *Loader) LoadPlugins() []driver.Plugin {
	var plugs = []driver.Plugin{}
	for _, p := range p.Plugins {
		plug, err := plugin.Open(p)
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		sym, err := plug.Lookup("Plugin")
		if err != nil {
			fmt.Println(err.Error())
			return nil
		}
		var myPlugin driver.Plugin
		myPlugin, ok := sym.(driver.Plugin)
		if !ok {
			grpclog.Fatalf("provided plugin: %T does not satisfy Plugin interface\n", sym)
			return nil
		} else {
			grpclog.Infof("registered plugin: %T\n", sym)
			plugs = append(plugs, myPlugin)
		}
	}

	return plugs
}
