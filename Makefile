build:
	go build -o bin/app ./cmd

run:
	./bin/app

test:
	go test -v ./...
