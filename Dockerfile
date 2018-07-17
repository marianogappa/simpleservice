FROM golang:latest
EXPOSE 8080
WORKDIR /go/src/github.com/marianogappa/simpleservice
RUN go get github.com/satori/go.uuid && \
    go get github.com/mesg-foundation/core/api/service && \
    go get google.golang.org/grpc
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o main .
CMD ["/go/src/github.com/marianogappa/simpleservice/main"]
