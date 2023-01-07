
FROM golang:1.19.1 

WORKDIR /app
ENV MOD=$GOPATH/pkg/mod
ENV GO111MODULE=on
RUN go install github.com/cucumber/godog/cmd/godog@latest
RUN go mod init
RUN go mod tidy


