package config

import (
	"crypto/tls"
	"fmt"
	"github.com/autom8ter/engine/driver"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"net"
	"net/http"
	"os"
	"path/filepath"
	pkg_runtime "runtime"
	"time"

	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"log"
)

func init() {
	fmt.Println("loading configuration settings with defaults")
	viper.AutomaticEnv()
	viper.AddConfigPath(".")
	viper.AddConfigPath("config")
	viper.AddConfigPath("plugins")
	{
		viper.SetDefault("grpc", &Address{
			Network: "tcp",
			Addr:    ":3000",
		})
		viper.SetDefault("gateway", &Address{
			Network: "tcp",
			Addr:    ":3000",
		})
		if pkg_runtime.GOOS == "windows" {
			viper.SetDefault("internal", &Address{
				Network: "tcp",
				Addr:    ":5050",
			})
		} else {
			viper.SetDefault("internal", &Address{
				Network: "unix",
				Addr:    "tmp/server.sock",
			})
		}
	}
	if err := viper.ReadInConfig(); err != nil {
		log.Println(err.Error())
	} else {
		fmt.Printf("using config file: %s\n", viper.ConfigFileUsed())
	}
}

// Config contains configurations of gRPC and Gateway server.
type Config struct {
	GrpcAddr                        *Address `mapstructure:"grpc" json:"grpc"`
	GrpcInternalAddr                *Address `mapstructure:"internal" json:"internal"`
	GatewayAddr                     *Address `mapstructure:"gateway" json:"gateway"`
	Swagger                         string   `mapstructure:"swagger" json:"swagger"`
	Plugins                         []driver.Plugin
	GrpcServerUnaryInterceptors     []grpc.UnaryServerInterceptor
	GrpcServerStreamInterceptors    []grpc.StreamServerInterceptor
	GatewayServerUnaryInterceptors  []grpc.UnaryClientInterceptor
	GatewayServerStreamInterceptors []grpc.StreamClientInterceptor
	GrpcServerOption                []grpc.ServerOption
	GatewayDialOption               []grpc.DialOption
	GatewayMuxOptions               []runtime.ServeMuxOption
	GatewayServerConfig             *HTTPServerConfig
	MaxConcurrentStreams            uint32
	GatewayServerMiddlewares        []driver.HTTPServerMiddleware
}

func New(plugins ...driver.Plugin) *Config {
	internal, ok := configAddr("internal")
	if !ok {
		internal = &Address{
			Network: "unix",
			Addr:    "tmp/server.sock",
		}
	}
	if internal.Network == "" {
		internal.Network = "unix"
	}
	if internal.Addr == "" {
		internal.Addr = "tmp/server.sock"
	}
	gateway, ok := configAddr("grpc")
	if !ok {
		gateway = &Address{
			Network: "tcp",
			Addr:    ":3000",
		}
	}
	if gateway.Network == "" {
		internal.Network = "tcp"
	}
	if gateway.Addr == "" {
		internal.Addr = ":3000"
	}
	return &Config{
		GrpcInternalAddr: internal,
		GatewayAddr:      gateway,
		GatewayServerConfig: &HTTPServerConfig{
			ReadTimeout:  8 * time.Second,
			WriteTimeout: 8 * time.Second,
			IdleTimeout:  2 * time.Minute,
		},
		MaxConcurrentStreams: 1000,
		Plugins:              plugins,
	}
}

// Address represents a network end point address.
type Address struct {
	Network string `mapstructure:"network" json:"network"`
	Addr    string `mapstructure:"addr" json:"addr"`
}

func (a *Address) CreateListener() (net.Listener, error) {
	if a.Network == "unix" {
		dir := filepath.Dir(a.Addr)
		f, err := os.Stat(dir)
		if err != nil {
			if err = os.MkdirAll(dir, 0755); err != nil {
				return nil, errors.Wrap(err, "failed to create the directory")
			}
		} else if !f.IsDir() {
			return nil, errors.Errorf("file %q already exists", dir)
		}
	}
	lis, err := net.Listen(a.Network, a.Addr)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to listen %s %s", a.Network, a.Addr)
	}
	return lis, nil
}

type HTTPServerConfig struct {
	TLSConfig         *tls.Config
	ReadTimeout       time.Duration `mapstructure:"read_timeout" json:"read_timeout"`
	ReadHeaderTimeout time.Duration `mapstructure:"read_header_timeout" json:"read_header_timeout"`
	WriteTimeout      time.Duration `mapstructure:"write_timeout" json:"write_timeout"`
	IdleTimeout       time.Duration `mapstructure:"idle_timeout" json:"idle_timeout"`
	MaxHeaderBytes    int           `mapstructure:"max_header_bytes" json:"max_header_bytes"`
	TLSNextProto      map[string]func(*http.Server, *tls.Conn, http.Handler)
	ConnState         func(net.Conn, http.ConnState)
}

func (c *HTTPServerConfig) InitConfig() *HTTPServerConfig {
	if err := viper.Unmarshal(c); err != nil {
		log.Fatal(err.Error())
	}
	return c
}

func (c *HTTPServerConfig) ApplyTo(s *http.Server) {
	s.TLSConfig = c.TLSConfig
	s.ReadTimeout = c.ReadTimeout
	s.ReadHeaderTimeout = c.ReadHeaderTimeout
	s.WriteTimeout = c.WriteTimeout
	s.IdleTimeout = c.IdleTimeout
	s.MaxHeaderBytes = c.MaxHeaderBytes
	s.TLSNextProto = c.TLSNextProto
	s.ConnState = c.ConnState
}

func (c *Config) ServerOptions() []grpc.ServerOption {
	return append(
		[]grpc.ServerOption{
			grpc_middleware.WithUnaryServerChain(c.GrpcServerUnaryInterceptors...),
			grpc_middleware.WithStreamServerChain(c.GrpcServerStreamInterceptors...),
			grpc.MaxConcurrentStreams(c.MaxConcurrentStreams),
		},
		c.GrpcServerOption...,
	)
}

func (c *Config) ClientOptions() []grpc.DialOption {
	return append(
		[]grpc.DialOption{
			grpc.WithInsecure(),
			grpc.WithDialer(func(a string, t time.Duration) (net.Conn, error) {
				return net.Dial(c.GrpcInternalAddr.Network, a)
			}),
			grpc.WithUnaryInterceptor(
				grpc_middleware.ChainUnaryClient(c.GatewayServerUnaryInterceptors...),
			),
			grpc.WithStreamInterceptor(
				grpc_middleware.ChainStreamClient(c.GatewayServerStreamInterceptors...),
			),
		},
		c.GatewayDialOption...,
	)
}

func Debug() {
	viper.Debug()
}

func Settings() map[string]interface{} {
	return viper.AllSettings()
}

func SetEnvPrefix(prefix string) {
	viper.SetEnvPrefix(prefix)
}

func BindFlags(cmd *cobra.Command) {
	_ = viper.BindPFlags(cmd.PersistentFlags())
	_ = viper.BindPFlags(cmd.Flags())
	for _, c := range cmd.Commands() {
		_ = viper.BindPFlags(c.PersistentFlags())
		_ = viper.BindPFlags(c.Flags())
	}
}

func configAddr(key string) (*Address, bool) {
	addr := viper.Get(key)
	switch addr.(type) {
	case *Address:
		return addr.(*Address), true
	default:
		return nil, false
	}
}
