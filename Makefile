.PHONY: help
help:	## show this help
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

.PHONY: check
check:	## go format ./..., go vet ./..., go install ./..., git add ., git commit -m "check"
	@go fmt ./...
	@go vet ./...
	@go test ./...
	@go install ./...
	@git add .
	@git commit -m "check"

build:
	docker build -t enginectl .

run:
	docker run --name enginectl -t -d enginectl

prune:
	docker container prune