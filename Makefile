BIN=sebas

.PHONY: install test desktop

all: install

install:
	go get ./...

test: install
	go test ./...

desktop: install
	(cd cmd/desktop && go run main.go)
