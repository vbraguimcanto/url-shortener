run:
	@go run main.go

docker-run:
	@docker compose build && docker compose up

test:
	@go test -v ./...