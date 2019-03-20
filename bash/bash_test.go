package bash_test

import (
	"fmt"
	"github.com/autom8ter/engine/bash"
	"testing"
)

func TestBash(t *testing.T) {
	bits, err := bash.Bash(`echo "hello world"`)
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(fmt.Sprintf("%s", bits))
}
