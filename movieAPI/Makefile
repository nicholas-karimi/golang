
BINARY_NAME=movieAPI
.DEFAULT_GOAL := run

build:
	# go build -o bin/movieAPI
	GOARCH=amd64 GOOS=linux go build -o ./target/${BINARY_NAME}-linux main.go
	GOARCH=amd64 GOOS=windows go build -o ./target/${BINARY_NAME}-windows main.go
	GOARCH=amd64 GOOS=darwin go build -o ./target/${BINARY_NAME}-darwin main.go

run: build
	# ./bin/movieAPI
	./target/${BINARY_NAME}-linux

clean:
	go clean
	rm ./bin/movieAPI
	rm ./target/${BINARY_NAME}-linux
	rm ./target/${BINARY_NAME}-windows
	rm ./target/${BINARY_NAME}-darwin

test:
	go test -v ./...


dep: 
	go mod download

vet:
	go mod vet

lint:
	golangci-lint run --enable-all
