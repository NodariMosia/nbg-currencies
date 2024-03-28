echo:
	echo "Hello World!"

setup:
	go version
	go mod tidy
	go mod download

run:
	go run ./cmd

clean:
	go clean
	rimraf bin

build:
	go mod tidy
	go build -o bin/ ./cmd

launch:
	./bin/cmd
