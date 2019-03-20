.PHONY: help
help:	## show this help
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

.PHONY: check
check:	## go format ./..., go vet ./..., go install ./..., git add ., git commit -m "check"
	@go fmt ./...
	@go vet ./...
	@go install ./...
	@git add .

build:
	docker build -t enginectl .

run:
	docker run --name enginectl -d -p 3000:3000 colemanword/enginectl init

prune:
	docker container stop enginectl
	docker container prune

.PHONY: version
version:	## go format ./..., go vet ./..., go install ./..., git add ., git commit -m "check"
	@go fmt ./...
	@go vet ./...
	@go install ./...
	@git add .
	@git commit -m "version"