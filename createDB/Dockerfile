ARG CODE_VERSION=latest
FROM golang:${CODE_VERSION} AS golang
ARG CODE_VERSION

LABEL goal="Download files using api-rest across go"

RUN ["/bin/sh", "-c", "echo my version of golang is the ${CODE_VERSION}"]

WORKDIR /my/app
COPY ./src ./

ENTRYPOINT [ "./entrypoint.sh" ]