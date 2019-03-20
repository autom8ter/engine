# Engine

`go get github.com/autom8ter/engine`

### Example

```go
package engine_test

import (
	"github.com/autom8ter/engine"
	"github.com/autom8ter/engine/testing/helloworld"
	"testing"
)

func TestNewEngine(t *testing.T) {
	//Create new engine instance from one or more driver.Plugins-> Register config options-> start server
	if err := engine.New(helloworld.NewBasicGreeter()).With(engine.WithDefaultLogger()).Serve(); err != nil {
		t.Fatal(err.Error())
	}
}
```

## ProtoCtl
`go get github.com/autom8ter/engine/protoctl`

### Example
