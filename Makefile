.PHONY: check
check:	## go format ./..., go vet ./..., go install ./..., git add ., git commit -m "check"
	@go fmt ./...
	@go vet ./...
	@go install ./...
	@go generate ./...
	@git add .
	@git commit -m "pass ✅"

clean: ## rm bin/*
	rm bin/*

.PHONY: help
help:	## show this help
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'