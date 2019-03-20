package plugin

import (
	"errors"
	"fmt"
	"github.com/autom8ter/engine/driver"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"log"
	"os"
	"path/filepath"
	"plugin"
)

func init() {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		home = os.Getenv("HOME")
	}
	fmt.Println("registered home directory: "+home)
	pluginPath=home+"/.plugins"
	fmt.Println("registered plugin path: "+pluginPath)
}

var pluginPath string

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
	fmt.Println(fmt.Sprintf("registered plugin: %T", sym))
	return myPlugin
}

func Files() []string {
	var files = []string{}
	if err := filepath.Walk(pluginPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}
		files = append(files, path)
		return nil
	}); err != nil {
		log.Fatal(err.Error())
	}
	return files
}

func LoadPlugins() []driver.Plugin {
	var files = Files()
	var plugs = []driver.Plugin{}
	for _, path := range files {
		newPlug := NewPluginLoader(path, "Plugin")
		plugs = append(plugs, newPlug.AsPlugin())
	}
	viper.Set("plugins", plugs)
	return plugs
}
