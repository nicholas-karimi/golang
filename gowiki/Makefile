BINARY_NAME=gowiki
.DEFAULT_GOAL := run
build:
	GOARCH=amd64 GOOS=linux go build -o ./target/${BINARY_NAME}-linux wiki.go	

run: build
		./target/${BINARY_NAME}-linux

clean:
	go clean
	rm ./target/${BINARY_NAME}-linux


test:
	go test -v ./...


dep:
	go mod download

vet:
	go mod vet

lint:
	golangcli-lint run --enable-all
