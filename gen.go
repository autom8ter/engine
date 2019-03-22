//go:generate env GOOS=linux go build -o bin/enginectl github.com/autom8ter/engine/enginectl
//go:generate env GOOS=darwin go build -o bin/enginectl_darwin github.com/autom8ter/engine/enginectl

package engine
