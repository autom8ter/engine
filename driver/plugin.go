package driver

import (
	"errors"
	"fmt"
	"plugin"
)

type PluginLoader struct {
	Path   string
	Symbol string
}

func NewPluginLoader(path string, symbol string) *PluginLoader {
	return &PluginLoader{Path: path, Symbol: symbol}
}

func (p *PluginLoader) AsPlugin() Plugin {
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
	var myPlugin Plugin
	myPlugin, ok := sym.(Plugin)
	if !ok {
		fmt.Println(errors.New("provided plugin does not satisfy Plugin interface"))
		return nil
	}
	return myPlugin
}
