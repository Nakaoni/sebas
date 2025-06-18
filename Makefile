BIN=sebas
APP_NAME=sebas

.PHONY: install test gui gui_dbg package

all: gui

install:
	go get ./...
	go install fyne.io/tools/cmd/fyne@latest

test: install
	go test ./...

gui: install
	(cd cmd/desktop && go run .)

gui_dbg: install
	(cd cmd/desktop && go run --tags debug .)

package: install
	(cd cmd/desktop && fyne package)
