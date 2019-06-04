ARG CODE_VERSION=latest
FROM golang:${CODE_VERSION} AS golang
ARG CODE_VERSION

RUN ["/bin/sh", "-c", "echo We are using the ${CODE_VERSION} of golang"]

WORKDIR $GOPATH/src
COPY src ./

RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o app .

EXPOSE 8080

CMD ["./app"]

#FROM alpine:${CODE_VERSION}
#ARG CODE_VERSION

#RUN ["/bin/sh", "-c", "echo We are using the ${CODE_VERSION} of alpine"]

#RUN apk --no-cache add ca-certificates

#WORKDIR /my
#COPY --from=golang /go/src/ .

#RUN ["/bin/sh", "-c", "ls -al"]
#RUN ["/bin/sh", "-c", "chmod +x app"]

#EXPOSE 8080

#CMD ["/bin/sh", "-c", "./app"]