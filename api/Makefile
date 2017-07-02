PROTOC :=$(shell command -v protoc 2> /dev/null)

build:
ifdef PROTOC
	go generate
endif

clean:
	-@rm -rf *.pb.go 2>/dev/null || true
