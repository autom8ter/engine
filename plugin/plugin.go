package plugin

import (
	"fmt"
	"github.com/autom8ter/engine/bash"
	"github.com/autom8ter/engine/driver"
	"github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
	"google.golang.org/grpc/grpclog"
	"log"
	"os"
	"path/filepath"
	"plugin"
	"strings"
)

func init() {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		home = os.Getenv("HOME")
	}
	pluginPath = home + "/.plugins"
	grpclog.Infof("registered plugin path: %s\n", pluginPath)
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
		dir, _ := filepath.Split(path)
		if strings.Contains(dir, "git") {
			return nil
		}
		if filepath.Ext(path) == ".git" {
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

func Build(file string) ([]byte, error) {
	return bash.Bash(GetScript(file))
}

func GetScript(file string) string {
	_, f := filepath.Split(file)
	fileStrip := strings.TrimSuffix(f, ".go")
	return fmt.Sprintf("env GOOS=linux go build -buildmode=plugin -o %s/%s.plugin %s", pluginPath, fileStrip, file)
}
