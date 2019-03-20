package main_test

import (
	"fmt"
	"github.com/autom8ter/engine/bash"
	"testing"
)

func TestBuild(t *testing.T) {
	script := "enginectl build -f ../testing/examplepb/example.go"
	bits, err := bash.Bash(script)
	if err != nil {
		t.Fatal(err.Error())
	}
	fmt.Println(string(bits))
}
