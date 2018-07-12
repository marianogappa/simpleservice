FROM golang:latest
EXPOSE 8080
WORKDIR /go/src/github.com/marianogappa/simpleservice
COPY . .
RUN GOOS=linux GOARCH=amd64 go build -o main .
CMD ["/go/src/github.com/marianogappa/simpleservice/main"]
