PROJECT_ROOT ?= $(shell pwd)
PROJECT_GO_BIN := $(PROJECT_ROOT)/go-bin

GOLANGCI_LINT_VERSION := v1.54.2

RUN_GOLANGCI_LINT = cd "$(PROJECT_ROOT)" && $(PROJECT_GO_BIN)/golangci-lint run -v --timeout 3m

default: lint test

lint:
	@$(RUN_GOLANGCI_LINT)

lint_fix:
	@$(RUN_GOLANGCI_LINT) --fix

test:
	go test -v

install: install_go_packages install_golangci

install_go_packages:
	@echo "$(shell go version)"
	@echo "GOPATH: $(GOPATH)"
	go get -u github.com/stretchr/testify

install_golangci:
	mkdir -p "$(PROJECT_GO_BIN)"
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b "$(PROJECT_GO_BIN)" $(GOLANGCI_LINT_VERSION)

.PHONY: default lint test install
