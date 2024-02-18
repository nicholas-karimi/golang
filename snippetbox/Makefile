BINARY_NAME=snippetbox
.DEFAULT_GOAL :=run

build:
	GOARCH=amd64 GOOS=linux go build -o ./target/${BINARY_NAME}-linux cmd/web/main.go

run: build
	./target/${BINARY_NAME}-linux

clean:
	go clean
	rm ./target/${BINARY_NAME}-linux

test:
	go test -v ./...

.PHONY: build run clean test