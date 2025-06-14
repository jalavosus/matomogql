GO = $(shell which go)

PROJECT = matomogql
BIN = server

CMD_DIR = ./cmd
BIN_DIR = ./bin
OUT_DIR = ./out

.PHONY: all build lint test clean generate fmt

all : clean build

clean :
	rm -rf $(BIN_DIR)/$(BIN)

build :
	$(GO) build -o $(BIN_DIR)/$(BIN) $(CMD_DIR)/$(BIN)

lint :
	golangci-lint run -c .golangci.yml

test :
	$(GO) test ./...

generate : # just add $(GO) generate [target] as needed
	$(GO) generate graph/resolver.go

fmt:
	$(GO) fmt ./...