.PHONY: check
check:	## go format ./..., go vet ./..., go install ./..., git add ., git commit -m "check"
	@go fmt ./...
	@go vet ./...
	@go install ./...
	@go generate ./...
	@go test -coverprofile COVERAGE.txt ./...
	@git add .
	@git commit -m "pass ✅"

build: ## docker build -t enginectl .
	docker build -t enginectl .

run: ## docker run --name enginectl -d -p 3000:3000 colemanword/enginectl serve
	docker run --name enginectl -d -p 3000:3000 colemanword/enginectl serve

prune: ## stop enginectl container, then prune all stopeed containers
	docker container stop enginectl
	docker container prune

clean: ## rm bin/*
	rm bin/*

.PHONY: help
help:	## show this help
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'