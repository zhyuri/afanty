PROTOC :=$(shell command -v protoc 2> /dev/null)

build:
ifdef PROTOC
		protoc --go_out=. *.proto
endif

clean:
	rm *.pb.go
