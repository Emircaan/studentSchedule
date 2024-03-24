build:
	go build -o bin/app ./cmd

run:

	./bin/app

run\:dev:

	wgo run  ./cmd

	

test:
	go test -v ./...
