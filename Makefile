## help: Get more info on available make commands.
help: Makefile
	@echo " Choose a command run in "$(PROJECTNAME)":"
	@sed -n 's/^##//p' $< | column -t -s ':' |  sed -e 's/^/ /'
.PHONY: help

## test: Run tests.
test:
	@echo "--> Running tests"
	@go test -timeout 5m ./...
.PHONY: test

## lint: Run golangci-lint.
lint:
	@echo "--> Running golangci-lint"
	@golangci-lint run
.PHONY: lint
