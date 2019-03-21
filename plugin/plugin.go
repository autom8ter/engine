package plugin

import (
	"fmt"
	"github.com/autom8ter/engine/driver"
	"github.com/spf13/viper"
	"google.golang.org/grpc/grpclog"
	"log"
	"os"
	"path/filepath"
	"plugin"
	"strings"
)

func init() {
	var err error
	pth, err := os.Getwd()
	if err != nil {
		pluginPath = os.Getenv("PWD") + "/plugins"
	} else {
		pluginPath = pth + "/plugins"
	}
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
		grpclog.Fatalf("provided plugin: %T does not satisfy Plugin interface\n", sym)
		return nil
	}
	grpclog.Infof("registered plugin: %T\n", sym)
	return myPlugin
}

func Files() []string {
	var files = []string{}
	if err := filepath.Walk(pluginPath, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if strings.Contains(path, "ingore") {
			return nil
		}
		if strings.Contains(path, "yaml") {
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
	viper.Set("files", files)
	return plugs
}
