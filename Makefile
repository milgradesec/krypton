VERSION:=$(shell git describe --tags --always --dirty="-dev")
BUILDFLAGS:=-v -ldflags="-s -w -X main.Version=$(VERSION)"
SYSTEM:=
IMPORT_PATH:=github.com/milgradesec/krypton

.PHONY: all
all: build

.PHONY: build
build:
	go build $(BUILDFLAGS) $(IMPORT_PATH)/cmd/krypton

.PHONY: release
release:
	set CGO_ENABLED=0
	set GOOS=windows
	set GOARCH=386
	go build $(BUILDFLAGS) $(IMPORT_PATH)

.PHONY: test
test:
	go test -v ./...

.PHONY: clean
clean:
	go clean