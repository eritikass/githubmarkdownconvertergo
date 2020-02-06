lint:
	$(GOPATH)/bin/golint .

install:
	go get -u golang.org/x/lint/golint

.PHONY: lint install
