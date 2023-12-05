include .env
export

# Tools.
TOOLS = ./tools
TOOLS_BIN = $(TOOLS)/bin

generate-swagger-auth:
	go generate ./...
	docker run --rm -it  \
		-u $(shell id -u):$(shell id -g) \
		-e GOPATH=$(shell go env GOPATH):/go \
		-e GOCACHE=/tmp \
		-v $(HOME):$(HOME) \
		-w $(shell pwd) \
		quay.io/goswagger/swagger:0.30.4 \
		generate spec -c ./cmd/auth --scan-models -c ./internal/auth -o ./swagger/OpenAPI/auth/rest.swagger.json


generate-auth-swagger:
	swagger generate spec --scan-models -c ./internal/auth -c ./cmd/auth -o ./swagger/OpenAPI/auth/rest.swagger.json

generate-swagger-pickup:
	go generate ./...
	docker run --rm -it  \
		-u $(shell id -u):$(shell id -g) \
		-e GOPATH=$(shell go env GOPATH):/go \
		-e GOCACHE=/tmp \
		-v $(HOME):$(HOME) \
		-w $(shell pwd) \
		quay.io/goswagger/swagger:0.30.4 \
		generate spec -c ./cmd/pickup --scan-models -c ./internal/pickup -o ./swagger/OpenAPI/pickup/rest.swagger.json

generate-pickup-swagger:
	swagger generate spec --scan-models -c ./internal/pickup -c ./cmd/pickup -o ./pickup/OpenAPI/pickup/rest.swagger.json

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


auth-migrate:
	migrate create -ext sql -dir migration/auth -seq create
	migrate create -ext sql -dir migration/auth -seq insert

auth-migrateup:
	migrate -path migration/auth -database "${AUTH_DATABASE_URL}" -verbose up

pickup-migrate:
	migrate create -ext sql -dir migration/pickup -seq create
	migrate create -ext sql -dir migration/pickup -seq insert

pickup-migrateup:
	migrate -path migration/pickup -database "${PICKUP_DATABASE_URL}" -verbose up

user-migrate:
	migrate create -ext sql -dir migration/user -seq create
	migrate create -ext sql -dir migration/user -seq insert

user-migrateup:
	migrate -path migration/user -database ${USER_DATABASE_URL} -verbose up

# sorting: sort_order=val&sort_by=asc/desc
# searching: field=val
# filtering:
#			num_field=lt/gt/eq/lte/gte:val, val:val
#         	date_field=val, val:val
#			bool_field=val
#           string_field=val