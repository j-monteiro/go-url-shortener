BIN_NAME=url_shortener.out

# Build the binary
build:
	go build -o bin/$(BIN_NAME) main.go

# Run the binary
run:
	go run main.go

test:
	go test -v ./...
