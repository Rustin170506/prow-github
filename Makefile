PROJECT=go-boilerplate
GOPATH ?= $(shell go env GOPATH)
P=8

# Ensure GOPATH is set before running build process.
ifeq "$(GOPATH)" ""
  $(error Please set the environment variable GOPATH before running `make`)
endif
FAIL_ON_STDOUT := awk '{ print } END { if (NR > 0) { exit 1 } }'

CURDIR := $(shell pwd)
path_to_add := $(addsuffix /bin,$(subst :,/bin:,$(GOPATH))):$(PWD)/tools/bin
export PATH := $(path_to_add):$(PATH)

GO              := GO111MODULE=on go
GOBUILD         := $(GO) build
GOTEST          := $(GO) test -p $(P)

PACKAGE_LIST  := go list ./...
PACKAGES  := $$($(PACKAGE_LIST))
PACKAGE_DIRECTORIES := $(PACKAGE_LIST) | sed 's|github.com/tidb-community-bots/$(PROJECT)||'
FILES     := $$(find .$$($(PACKAGE_DIRECTORIES)) -name "*.go")


.PHONY: build clean test dev check tidy

build:
	$(GOBUILD)

clean:
	$(GO) clean -i ./...
	rm -rf *.out

test:
	$(GOTEST)
	@>&2 echo "Great, all tests passed."

dev: check test

check: fmt tidy staticcheck

fmt:
	@echo "gofmt (simplify)"
	@gofmt -s -l -w $(FILES) 2>&1 | $(FAIL_ON_STDOUT)

tidy:
	@echo "go mod tidy"
	./tools/check/check-tidy.sh

staticcheck: tools/bin/golangci-lint
	tools/bin/golangci-lint run  $$($(PACKAGE_DIRECTORIES))

tools/bin/golangci-lint:
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s -- -b ./tools/bin v1.31.0

