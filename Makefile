build:
	go build

run:
	go run .

test:
	go test ./...

test-integration:
	go test ./... --tags=integration

run-docker:
	docker build -t fizzbuzz --pull . && docker run -it --rm fizzbuzz

.PHONY: build run test test-integration run-docker