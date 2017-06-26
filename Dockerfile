FROM golang:1.8 as builder

RUN mkdir -p /go/src/github.com/zhyuri/afanty
COPY . /go/src/github.com/zhyuri/afanty
RUN go install github.com/zhyuri/afanty

ENTRYPOINT ["/go/bin/afanty"]

FROM alpine:latest
COPY --from=builder /go/bin/afanty /bin/afanty
ENTRYPOINT ["/bin/afanty"]
