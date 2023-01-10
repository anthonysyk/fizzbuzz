build:
	go build

run:
	go run .

test:
	go test ./...

docker:
	docker build -t fizzbuzz --pull . && docker run -it --rm fizzbuzz
