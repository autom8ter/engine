package plugin_test

import (
	"fmt"
	"github.com/autom8ter/engine/plugin"
	"testing"
)

func TestBuildCmd(t *testing.T) {
	fmt.Println("script: ", plugin.GetScript("../testing/examplepb/example.go"))
	bits, err := plugin.Build("../testing/examplepb/example.go")
	if err != nil {
		t.Fatalf("cmd.Run() failed: \n%s\n", err)
	}
	fmt.Println(string(bits))

}
