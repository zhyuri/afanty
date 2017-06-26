FROM golang:1.8 as builder

RUN go get -u github.com/golang/dep/cmd/dep && \
mkdir -p /go/src/github.com/zhyuri/afanty
COPY . /go/src/github.com/zhyuri/afanty

WORKDIR /go/src/github.com/zhyuri/afanty
RUN dep ensure && go install github.com/zhyuri/afanty

ENTRYPOINT ["/go/bin/afanty"]

FROM alpine:latest
COPY --from=builder /go/bin/afanty /bin/afanty
ENTRYPOINT ["/bin/afanty"]
