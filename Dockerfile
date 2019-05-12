ARG CODE_VERSION=latest
FROM golang:${CODE_VERSION} AS golang_and_go
ARG CODE_VERSION

RUN ["/bin/sh", "-c", "echo we are running the ${CODE_VERSION} version of golang"]

WORKDIR /my/app
COPY ./src ./

ENTRYPOINT ["./run.sh"]