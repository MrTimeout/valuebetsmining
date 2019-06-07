ARG CODE_VERSION=latest
FROM golang:${CODE_VERSION} AS golang
ARG CODE_VERSION

RUN ["/bin/sh", "-c", "echo We are using the ${CODE_VERSION} of golang"]

WORKDIR $GOPATH/src
COPY data ./

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o app .

EXPOSE 8080

CMD ["./app"]