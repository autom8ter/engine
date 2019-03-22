//go:generate env GOOS=linux go build -o release/enginectl github.com/autom8ter/engine/enginectl
//go:generate env GOOS=darwin go build -o release/enginectl_darwin github.com/autom8ter/engine/enginectl
//go:generate godocdown -o GODOC.md
//go:generate go test -coverprofile COVERAGE.txt ./...
//go:generate go tool cover -html=COVERAGE.txt -o COVERAGE.html

package engine
