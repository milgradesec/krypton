VERSION:=$(shell git describe --tags --always --abbrev=0 --dirty="-dev")
BUILDFLAGS:=-v -ldflags="-s -w"
IMPORT_PATH:=github.com/milgradesec/krypton

.PHONY: all
all: build

.PHONY: build
build:
	go build $(BUILDFLAGS)

.PHONY: test
test:
	go test -v ./...

.PHONY: lint
lint:
	golangci-lint run

.PHONY: clean
clean:
	go clean