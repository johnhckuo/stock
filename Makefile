docker:
	docker build -f ./deploy/Dockerfile -t="codementordev/hello-johnhckuo" .

test-api:
	go test -v ./test/ -count 1

run:
	go run cmd/codementordev-hello-johnhckuo-server/main.go --port=8080

lint:
	golint ./...

fmt:
	go fmt ./...

swagger:
	swagger generate server -A codementordev/hello-johnhckuo

docs:
	swagger generate markdown -f ./swagger.yml --output ./spec.md

	