docker:
	docker build -f ./deploy/Dockerfile -t="codementordev/hello-johnhckuo" .

run:
	go run cmd/codementordev-hello-johnhckuo-server/main.go --port=8080

swagger:
	swagger generate server -A codementordev/hello-johnhckuo

docs:
	swagger serve swagger.yml

spec:
	swagger generate markdown -f ./swagger.yml --output ./spec.md

unit_test:
	go test ./internal/adaptors/ -v -count=1 -failfast -coverprofile cover.out && go tool cover -html=internal/adaptors/cover.out

	