ARG CODE_VERSION=latest
FROM golang:${CODE_VERSION} AS golang
ARG CODE_VERSION

RUN ["/bin/sh", "-c", "echo We are using the $CODE_VERSION of golang"]

WORKDIR $GOPATH/src/valuebetsmining/data/
COPY data ./

RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -installsuffix cgo -o app .

FROM alpine:${CODE_VERSION} AS alpine
ARG CODE_VERSION

RUN ["/bin/sh", "-c", "echo We are using the ${CODE_VERSION} of alpine"]

RUN apk --no-cache add ca-certificates

WORKDIR /data

RUN mkdir -p leagues/CSV/
COPY --from=golang go/src/valuebetsmining/data/ .

RUN /data/app run --run-connection=true --run-output=true

FROM mongo:${CODE_VERSION}
ARG CODE_VERSION

#ENV 

RUN ["/bin/sh", "-c", "echo We are using the ${CODE_VERSION} of mongoDB"]

EXPOSE 27017

COPY --from=alpine /data/leagues /leagues
COPY ./data/script /docker-entrypoint-initdb.d
