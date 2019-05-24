ARG CODE_VERSION=latest
FROM golang:${CODE_VERSION} AS golang
ARG CODE_VERSION

RUN ["/bin/sh", "-c", "echo We are using the ${CODE_VERSION} of golang"]

WORKDIR $GOPATH/src
COPY ./src ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o data .

FROM alpine:${CODE_VERSION}
ARG CODE_VERSION

RUN ["/bin/sh", "-c", "echo We are using the ${CODE_VERSION} of alpine"]

RUN apk --no-cache add ca-certificates

WORKDIR /my
COPY --from=golang /go/src/ .

CMD ["./data"]