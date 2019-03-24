// Package config is used to setup the configuration of a new Engine instance. A basic config instance
// is created from your config file with New() and then it may be configured more with it's method With(...options)

package config

import (
	"github.com/autom8ter/engine/driver"
	"github.com/autom8ter/engine/util"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
	"os"
	"plugin"
)

func init() {
	grpclog.SetLoggerV2(grpclog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout))
}

// Config contains configurations of gRPC and Gateway server. A new instance of Config is created from your config.yaml|config.json file in your current working directory
// Network, Address, and Paths can be set in your config file to set the Config instance. Otherwise, defaults are set.
type Config struct {
	Network            string   `json:"network"`
	Address            string   `json:"address"`
	Paths              []string `json:"paths"`
	Symbol             string   `json:"symbol"`
	Plugins            []driver.Plugin
	UnaryInterceptors  []grpc.UnaryServerInterceptor
	StreamInterceptors []grpc.StreamServerInterceptor
	Option             []grpc.ServerOption
}

// New creates a config from your config file. If no config file is present, the resulting Config will have the following defaults: netowork: "tcp" address: ":3000"
// use the With method to continue to modify the resulting Config object
func New(network, addr string) *Config {
	if network == "" || addr == "" {
		util.Debugf("empty network or address detected %s %s, setting defaults: tcp :3000\n", network, addr)
		network = "tcp"
		addr = ":3000"
	}
	c := &Config{
		Network: network,
		Address: addr,
		Symbol:  "Plugin",
	}
	return c
}

// CreateListener creates a network listener from the network and address config
func (c *Config) CreateListener() (net.Listener, error) {
	util.Debugf("creating server listener %s %s\n", c.Network, c.Address)
	lis, err := net.Listen(c.Network, c.Address)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to listen %s %s", c.Network, c.Address)
	}
	return lis, nil
}

// With is used to configure/initialize a Config with custom options
func (c *Config) With(opts ...Option) *Config {
	for _, f := range opts {
		grpclog.Infoln("Adding option")
		f(c)
	}
	return c
}

// LoadPlugins loads driver.Plugins from paths set with config.WithPluginPaths(...)
func (c *Config) loadPlugins() {
	for _, p := range c.Paths {
		util.Debugf("registered path: %v\n", p)
		plug, err := plugin.Open(p)
		if err != nil {
			grpclog.Fatalln(err.Error())
		}
		sym, err := plug.Lookup(c.Symbol)
		if err != nil {
			grpclog.Fatalln(err.Error())
		}

		var asPlugin driver.Plugin
		asPlugin, ok := sym.(driver.Plugin)
		if !ok {
			grpclog.Fatalf("provided plugin: %T does not satisfy Plugin interface\n", sym)
		} else {
			util.Debugf("registered plugin: %T\n", sym)
			c.Plugins = append(c.Plugins, asPlugin)
		}
	}
	if len(c.Plugins) == 0 {
		grpclog.Warningln("No plugins detected. 0 registered plugins.")
	}
}