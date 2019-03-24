package main

import (
	"github.com/autom8ter/engine"
	"github.com/autom8ter/engine/config"
	"github.com/autom8ter/engine/examples/middleware/logging"
	"log"
)

func main() {
	e := engine.New("tcp", ":8080").With(
		config.WithDebug(),
		config.WithPlugins("Plugin", "bin/example.so"),
		config.WithChannelz(),
		config.WithReflection(),
		config.WithHealthz(),
		config.WithUnaryInterceptors(logging.NewUnaryLogger()),
	)
	defer e.Shutdown()
	switch {
	case e == nil:
		log.Fatalln("detected nil engine instance")
	case e.Config().Symbol != "Plugin":
		log.Fatalf("detected plugin symbol mismatched, got: %s\n", e.Config().Symbol)
	case len(e.Config().UnaryInterceptors) < 1:
		log.Fatalln("expected at least one unary interceptor")
	case e.Config().Network != "tcp":
		log.Fatalf("expected tcp network, got: %s\n", e.Config().Network)
	case e.Config().Address != ":8080":
		log.Fatalf("expected :8080 address, got: %s\n", e.Config().Address)
	}
	if err := e.Serve(); err != nil {
		log.Fatalln(err.Error())
	}
}
