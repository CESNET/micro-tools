SHELL := bash
NAME := micro-tools
IMPORT := github.com/CESNET/$(NAME)
UNAME := $(shell uname -s)
GOBUILD ?= go build
PACKAGES ?= $(shell go list ./...)
SOURCES ?= $(shell find . -name "*.go" -type f)
GENERATE ?= $(PACKAGES)

TAGS ?=

LDFLAGS += -s -w
GCFLAGS += all=-N -l

.PHONY: all
all: build

.PHONY: sync
sync:
	go mod download

.PHONY: clean
clean:
	go clean -i ./...

.PHONY: fmt
fmt:
	gofmt -s -w $(SOURCES)

.PHONY: vet
vet:
	go vet $(PACKAGES)

.PHONY: staticcheck
staticcheck:
	go run honnef.co/go/tools/cmd/staticcheck -tags '$(TAGS)' $(PACKAGES)

.PHONY: lint
lint:
	for PKG in $(PACKAGES); do go run golang.org/x/lint/golint -set_exit_status $$PKG || exit 1; done;

.PHONY: generate
generate:
	go generate $(GENERATE)

.PHONY: changelog
changelog:
	go run github.com/restic/calens

.PHONY: test
test:
	go run github.com/haya14busa/goverage -v -coverprofile coverage.out $(PACKAGES)

.PHONY: build
build:
	$(GOBUILD) -v -tags '$(TAGS)' -ldflags '$(LDFLAGS)' ./...
