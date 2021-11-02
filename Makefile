VERSION = $(file < VERSION)
COMMIT = $(shell git rev-parse HEAD)

all: drc

drc: go.mod go.sum $(shell find . -name '*.go') VERSION
	go build -o drc -ldflags "-X 'main.version=${VERSION}'"