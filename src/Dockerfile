ARG CODE_VERSION=latest
FROM golang:${CODE_VERSION} AS golang
ARG CODE_VERSION

RUN ["/bin/sh", "-c", "echo We are using the ${CODE_VERSION} of golang"]

WORKDIR $GOPATH/src/valuebetsmining/src/
COPY src ./

RUN go get -u github.com/golang/dep/cmd/dep

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o app .

EXPOSE 8080

CMD ["./app"]