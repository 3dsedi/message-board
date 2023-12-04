BINARY_NAME=bin/gobank

test:
	go test ./tests/...

build:
	@go build -o  ${BINARY_NAME}

run: build
	@./${BINARY_NAME}

clean:
	go clean
	rm ${BINARY_NAME}

test:
	@go test -v ./...
