ARG CODE_VERSION=latest
FROM golang:${CODE_VERSION} AS golang
ARG CODE_VERSION

RUN ["/bin/sh", "-c", "echo We are using the ${CODE_VERSION} of golang"]

WORKDIR $GOPATH/src/valuebetsmining/src/
COPY . .

RUN go get -u github.com/golang/dep/cmd/dep

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -installsuffix cgo -o app .

FROM alpine:${CODE_VERSION}
ARG CODE_VERSION

RUN ["/bin/sh", "-c", "echo We are using the ${CODE_VERSION} of alpine"]

RUN apk --no-cache add ca-certificates

#ENV 

WORKDIR /my
COPY --from=golang go/src/valuebetsmining/src/ .

EXPOSE 8080

CMD ["/my/app"]