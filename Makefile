docker:
	docker build -f ./deploy/Dockerfile -t="codementordev/hello-johnhckuo" .

test-api:
	go test -v ./test/ -count 1

run:
	go run cmd/server/main.go

lint:
	golint ./...

fmt:
	go fmt ./...
