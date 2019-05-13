#Golang compiling program
ARG CODE_VERSION=latest
FROM golang:${CODE_VERSION} AS golang_and_go
ARG CODE_VERSION

#Copy project
WORKDIR $GOPATH/src/valuebetsmining/src/
COPY ./src/ .

#Download dependencies and install them
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure

#Build
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

# Alpine deploying service
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /my
COPY --from=golang_and_go /go/src/valuebetsmining/src/ .

EXPOSE 80

CMD ["./app"]