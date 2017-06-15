DIST=dist/
BINARY=afanty

GIT_TAG=$(shell git rev-parse --short HEAD)
BUILD_TIME=$(shell date +'%Y-%m-%d %H:%M')

LDFLAGS=-ldflags '-X "main.gitTag=${GIT_TAG}" -X "main.buildTime=${BUILD_TIME}"'

build:
	dep ensure
	go build ${LDFLAGS} -o ${DIST}${BINARY} main.go

test:
	go test $(go list ./... | grep -v /vendor/)

run: build
	${DIST}${BINARY}

pack: build
	tar -cf ${BINARY}.tar dist
	rm -rf dist