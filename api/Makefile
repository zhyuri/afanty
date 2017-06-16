PROTOC :=$(shell command -v protoc 2> /dev/null)

build:
ifdef PROTOC
	protoc -I . \
		service.proto \
		--go_out=plugins=grpc:.
endif

clean:
	rm *.pb.go
