package plugin

import (
	"errors"
	"fmt"
	"github.com/autom8ter/engine/driver"
	"log"
	"plugin"
)

type PluginLoader struct {
	Path   string
	Symbol string
}

func NewPluginLoader(path string, symbol string) PluginLoader {
	return PluginLoader{Path: path, Symbol: symbol}
}

func (p PluginLoader) AsPlugin() driver.Plugin {
	plug, err := plugin.Open(p.Path)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	sym, err := plug.Lookup(p.Symbol)
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}
	var myPlugin driver.Plugin
	myPlugin, ok := sym.(driver.Plugin)
	if !ok {
		log.Fatalln(errors.New(fmt.Sprintf("provided plugin: %T does not satisfy Plugin interface", sym)))
		return nil
	}
	return myPlugin
}
