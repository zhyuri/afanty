DIST=dist/
BINARY=afanty

GIT_TAG=$(shell git rev-parse --short HEAD)
BUILD_TIME=$(shell date +'%Y-%m-%d %H:%M')

LDFLAGS=-ldflags '-X "main.gitTag=${GIT_TAG}" -X "main.buildTime=${BUILD_TIME}"'

build: clean pb
	dep ensure
	gcloud container builds submit -t asia.gcr.io/afanty-170802/afanty .

pb:
	$(MAKE) -C api

test:
	go test `go list ./... | grep -v /vendor/ | grep -v /api`

doc:
	godoc -http=:6060 -index

run:
	${DIST}${BINARY}

pull:
	gcloud docker -- pull asia.gcr.io/afanty-170802/afanty

clean:
	$(MAKE) clean -C api

.PHONY: pb build clean
