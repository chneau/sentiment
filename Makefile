.SILENT:
.ONESHELL:
.NOTPARALLEL:
.EXPORT_ALL_VARIABLES:
.PHONY: run deps build clean exec

run: buildPublic build exec clean

exec:
	./bin/app

buildPublic:
	statik -src=./static -dest=./pkg

build:
	go build -o bin/app -ldflags '-s -w -extldflags "-static"'

clean:
	rm -rf bin

deps:
	go get -d -u -v ./...

test:
	GOCACHE=off go test -v ./pkg/sentiment/...

testargv: buildPublic build
	./bin/app super happy
	$(MAKE) clean

