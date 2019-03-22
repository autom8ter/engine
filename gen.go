//go:generate env GOOS=linux go build -o release/enginectl github.com/autom8ter/engine/enginectl
//go:generate env GOOS=darwin go build -o release/enginectl_darwin github.com/autom8ter/engine/enginectl

package engine
