#Golang compiling program
ARG CODE_VERSION=latest
FROM golang:${CODE_VERSION} AS golang_and_go
ARG CODE_VERSION

WORKDIR /go/src/
COPY ./src/ ./

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Alpine deploying service
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /my
COPY --from=golang_and_go /my/app .

EXPOSE 80

CMD ["./app"]