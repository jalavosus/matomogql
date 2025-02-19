GO = $(shell which go)

PROJECT = matomogql
APP_SERVER = server

CMD_DIR = ./cmd
BIN_DIR = ./bin
OUT_DIR = ./out

.PHONY: all
all : clean

.PHONY: generate-schema
generate-schema :
	$(GO) generate graph/resolver.go

.PHONY: lint
lint :
	golangci-lint run -c .golangci.yml

.PHONY: test
test :
	$(GO) test ./...


.PHONY: clean
clean :
	rm -rf $(BIN_DIR)


.PHONY: build-server
build-server :
	$(GO) build -o $(BIN_DIR)/$(APP_SERVER) $(CMD_DIR)/$(APP_SERVER)