# Engine

Engine is a pluggable framework for rest/grpc services

`go get github.com/autom8ter/engine`

## API/Interface/Driver
/driver/driver.go
```go
package driver

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"net"
)

// Handler is an interface for representing gRPC server implementations.
type Plugin interface {
	RegisterWithServer(*grpc.Server)
	RegisterWithHandler(context.Context, *runtime.ServeMux, *grpc.ClientConn) error
}

// Server provides an interface for starting and stopping the server.
type Listener interface {
	Serve(l net.Listener) error
	Shutdown()
}

```

## Plugins

### Building

`go build -buildmode=plugin -o $PLUGIN.plugin $PLUGIN.go`

example: testing/examplepb

### Using

```go
package engine_test

import (
	"github.com/autom8ter/engine"
	"github.com/autom8ter/engine/config"
	"github.com/autom8ter/engine/plugin"
	"os"
	"testing"
)
var home = os.Getenv("HOME")

var examplePath = home+"/.plugins/example.plugin"

func TestNewEngine(t *testing.T) {
	if err := engine.New().With(config.WithPluginLoaders(plugin.PluginLoader{
		Path:   examplePath,
		Symbol: "ExamplePb",
	})).Serve(); err != nil {
		t.Fatal(err.Error())
	}
}

```