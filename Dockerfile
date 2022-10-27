FROM golang:1.18-alpine as builder

RUN mkdir /pwl-testapi
COPY go.mod /pwl-testapi/
COPY go.sum /pwl-testapi/
COPY main.go /pwl-testapi/

RUN cd /pwl-testapi;go mod tidy;go mod download;go build -o pwl-testapi
RUN echo "build complete"


#Runtime container
FROM alpine:3


COPY --from=builder /pwl-testapi/pwl-testapi /usr/local/bin/

RUN adduser -D certmanager
USER certmanager

ENTRYPOINT ["/usr/local/bin/pwl-testapi"]