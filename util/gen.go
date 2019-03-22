//go:generate godocdown -o README.md
//go:generate go test -coverprofile COVERAGE.txt ./...
//go:generate go tool cover -html=COVERAGE.txt -o COVERAGE.html

package util
