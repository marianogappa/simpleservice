all: build

build: test
	go build .

test:
	go test -v ./...
