default: lint test

lint:
	$(GOPATH)/bin/golint .

test:
	go test -v

install:
	@echo "$(shell go version)"
	@echo "GOPATH: $(GOPATH)"
	go get -u golang.org/x/lint/golint
	go get -u github.com/stretchr/testify

.PHONY: default lint test install
