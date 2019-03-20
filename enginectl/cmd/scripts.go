package cmd

import (
	"log"
	"os"
	"path/filepath"
)

func PluginPath() string {
	return os.Getenv("HOME") + "/.plugins"
}

func PluginsFound() []string {
	var files = []string{}
	if err := filepath.Walk(PluginPath(), func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	}); err != nil {
		log.Fatal(err.Error())
	}
	return files
}
