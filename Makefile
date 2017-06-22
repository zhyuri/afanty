DIST=dist/
BINARY=afanty

GIT_TAG=$(shell git rev-parse --short HEAD)
BUILD_TIME=$(shell date +'%Y-%m-%d %H:%M')

LDFLAGS=-ldflags '-X "main.gitTag=${GIT_TAG}" -X "main.buildTime=${BUILD_TIME}"'

build: clean pb
	dep ensure
	go build ${LDFLAGS} -o ${DIST}${BINARY} afanty.go

pb:
	$(MAKE) -C api

test:
	go test `go list ./... | grep -v /vendor/ | grep -v /api`

doc:
	godoc -http=:6060 -index

run:
	${DIST}${BINARY}

pack: build
	tar -cf ${BINARY}.tar dist
	rm -rf dist

clean:
	$(MAKE) clean -C api

.PHONY: pb build clean
