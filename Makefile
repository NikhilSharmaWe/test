build: 
	go build -o bin/test

run: build
	./bin/test

test:
	@go test -v ./...