FROM golang:1.14 as builder

RUN mkdir -p /go/src/github.com/hsmtkk/jissen_container

WORKDIR /go/src/github.com/hsmtkk/jissen_container

COPY ./pkg/smpwebsrv ./pkg/smpwebsrv
COPY ./cmd/smpwebsrv ./cmd/smpwebsrv
COPY ./go.mod .
COPY ./go.sum .

ENV CGO_ENABLED=0
RUN go build -o /smpwebsrv.bin

FROM alpine:3.11.5

COPY --from=builder /smpwebsrv.bin /smpwebsrv.bin

EXPOSE 80

ENTRYPOINT ["/smpwebsrv.bin"]
