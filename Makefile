set-env:
	cpp .env.example .env

build:
	@go build -C cmd/prod-ready-api -o ../../bin/prod-go

run: build
	@./bin/prod-go

test:
	@go test -v ./...

compose-local-up:
	docker compose -f docker-compose.yaml up -d