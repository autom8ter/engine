// Package config is used to setup the configuration of a new Engine instance. A basic config instance
// is created from your config file with New() and then it may be configured more with it's method With(...options)

package config

import (
	"github.com/autom8ter/engine/driver"
	"github.com/autom8ter/objectify"
	"github.com/autom8ter/util"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"net"
	"os"
)

var tool = objectify.Default()

// Config contains configurations of gRPC and Gateway server. A new instance of Config is created from your config.yaml|config.json file in your current working directory
// Network, Address, and Paths can be set in your config file to set the Config instance. Otherwise, defaults are set.
type Config struct {
	Network            string `json:"network" validate:"required"`
	Address            string `json:"address" validate:"required"`
	Plugins            []driver.Plugin
	UnaryInterceptors  []grpc.UnaryServerInterceptor
	StreamInterceptors []grpc.StreamServerInterceptor
	Option             []grpc.ServerOption
}

// New creates a config from your config file. If no config file is present, the resulting Config will have the following defaults: netowork: "tcp" address: ":3000"
// use the With method to continue to modify the resulting Config object
func New(network, addr string, debug bool) *Config {
	if debug {
		_ = os.Setenv("DEBUG", "t")
		_ = os.Setenv("debug", "t")
	}
	if network == "" || addr == "" {
		tool.Debugf("empty network or address detected, setting defaults %s", "tcp :3000")
		network = "tcp"
		addr = ":3000"
	}
	c := &Config{
		Network: network,
		Address: addr,
	}
	tool.PanicIfNil(c)
	if err := tool.Validate(c); err != nil {
		panic(errors.New("failed to valid config"))
	}
	return c
}

// CreateListener creates a network listener from the network and address config
func (c *Config) CreateListener() (net.Listener, error) {
	tool.PanicIfNil(c)
	lis, err := net.Listen(c.Network, c.Address)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to listen %s %s", c.Network, c.Address)
	}
	tool.PanicIfNil(lis)
	return lis, nil
}

// With is used to configure/initialize a Config with custom options
func (c *Config) With(opts ...Option) *Config {
	tool.PanicIfNil(c)
	for _, f := range opts {
		f(c)
	}
	return c
}

func (c *Config) ServerOptions() []grpc.ServerOption {
	tool.PanicIfNil(c)
	opts := append(
		[]grpc.ServerOption{
			grpc_middleware.WithUnaryServerChain(c.UnaryInterceptors...),
			grpc_middleware.WithStreamServerChain(c.StreamInterceptors...),
		},

		c.Option...,
	)
	return opts
}

func (c *Config) Debug() string {
	type cfgLog struct {
		Network            string `json:"network"`
		Address            string `json:"address"`
		UnaryInterceptors  int    `json:"unary_interceptors"`
		StreamInterceptors int    `json:"stream_interceptors"`
		Options            int    `json:"options"`
		Plugins            int    `json:"plugins"`
	}
	logcfg := &cfgLog{
		Network:            c.Network,
		Address:            c.Address,
		UnaryInterceptors:  len(c.UnaryInterceptors),
		StreamInterceptors: len(c.StreamInterceptors),
		Options:            len(c.ServerOptions()),
		Plugins:            len(c.Plugins),
	}
	return util.ToPrettyJsonString(logcfg)
}
