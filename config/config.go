// Package config is used to setup the configuration of a new Engine instance. A basic config instance
// is created from your config file with New() and then it may be configured more with it's method With(...options)

package config

import (
	"github.com/autom8ter/engine/driver"
	"github.com/autom8ter/engine/util"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	"net"
	"os"
)

func init() {
	grpclog.SetLoggerV2(grpclog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout))
	util.Debugln("initialized global logger")
}

// Config contains configurations of gRPC and Gateway server. A new instance of Config is created from your config.yaml|config.json file in your current working directory
// Network, Address, and Paths can be set in your config file to set the Config instance. Otherwise, defaults are set.
type Config struct {
	Network            string `json:"network"`
	Address            string `json:"address"`
	Plugins            []driver.Plugin
	UnaryInterceptors  []grpc.UnaryServerInterceptor
	StreamInterceptors []grpc.StreamServerInterceptor
	ServerOptions      []grpc.ServerOption
	RouterOptions      []RouterOption
	HTTPOptions        []HTTPOption
	match              cmux.CMux
}

// New creates a config from your config file. If no config file is present, the resulting Config will have the following defaults: netowork: "tcp" address: ":3000"
// use the With method to continue to modify the resulting Config object
func New(network, addr string) *Config {
	if network == "" || addr == "" {
		util.Debugf("empty network or address detected: %s %s, setting defaults: tcp, :3000\n", network, addr)
		network = "tcp"
		addr = ":3000"
	}
	util.Debugf("creating server listener %s %s\n", network, addr)
	lis, err := net.Listen(network, addr)
	if err != nil {
		panic(err.Error())
	}

	c := &Config{
		Network: network,
		Address: addr,
		match:   cmux.New(lis),
	}
	return c
}

// With is used to configure/initialize a Config with custom options
func (c *Config) With(opts ...Option) *Config {
	for _, f := range opts {
		grpclog.Infoln("Adding option")
		f(c)
	}
	return c
}

func (c *Config) HTTPHeaderField(name, value string) net.Listener {
	return c.match.Match(cmux.HTTP2HeaderField(name, value), cmux.HTTP1HeaderField(name, value))
}

func (c *Config) Any() net.Listener {
	return c.match.Match(cmux.Any())
}

func (c *Config) HTTP() net.Listener {
	return c.match.Match(cmux.HTTP1(), cmux.HTTP2())
}

func (c *Config) GRPC() net.Listener {
	return c.match.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
}

func (c *Config) Serve() error {
	return c.match.Serve()
}
