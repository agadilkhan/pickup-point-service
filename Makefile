# Tools.
TOOLS = ./tools
TOOLS_BIN = $(TOOLS)/bin

generate-swagger-user:
	go generate ./...
	docker run --rm -it  \
		-u $(shell id -u):$(shell id -g) \
		-e GOPATH=$(shell go env GOPATH):/go \
		-e GOCACHE=/tmp \
		-v $(HOME):$(HOME) \
		-w $(shell pwd) \
		quay.io/goswagger/swagger:0.30.4 \
		generate spec -c ./cmd/user --scan-models -c ./internal/user -o ./swagger/OpenAPI/user.rest.swagger.json

generate-user-swagger:
	swagger generate spec --scan-models -c ./internal/user -c ./cmd/user -o ./swagger/OpenAPI/user.rest.swagger.json


.PHONY: fix-lint
fix-lint: $(TOOLS_BIN)/golangci-lint
	$(TOOLS_BIN)/golangci-lint run --fix

imports: $(TOOLS_BIN)/goimports
	$(TOOLS_BIN)/goimports -local "service" -w ./internal ./cmd

# INSTALL linter
$(TOOLS_BIN)/golangci-lint: export GOBIN = $(shell pwd)/$(TOOLS_BIN)
$(TOOLS_BIN)/golangci-lint:
	mkdir -p $(TOOLS_BIN)
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.54.2


# INSTALL goimports
$(TOOLS_BIN)/goimports: export GOBIN = $(shell pwd)/$(TOOLS_BIN)
$(TOOLS_BIN)/goimports:
	mkdir -p $(TOOLS_BIN)
	go install golang.org/x/tools/cmd/goimports@latest

