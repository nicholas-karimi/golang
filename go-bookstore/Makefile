BINARY_NAME=bookstore
.DEFAULT_GOAL := run

build:
	GOARCH=amd64 GOOS=linux go build -o ./target/${BINARY_NAME}-linux main.go
	GOARCH=amd64 GOOS=darwin go build -o ./target/${BINARY_NAME}-darwin main.go
	GOARCH=amd64 GOOS=windows go build -o ./target/${BINARY_NAME}-windows main.go

run: build
	./target/${BINARY_NAME}-linux


clean:
	go clean

	rm ./target/${BINARY_NAME}-linux
	rm ./target/${BINARY_NAME}-darwin
	rm ./target/${BINARY_NAME}-windows

test:
	go test -v ./...

dep:
	go mod download

vet:
	go mod vet

lint:
	golangcli-lint run --enable-all
	
		
		
		
