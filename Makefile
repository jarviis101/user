include .env
export

GOOSE=go run github.com/pressly/goose/v3/cmd/goose@v3.27.1

migration-create:
	@read -p "Enter migration name: " name; \
	$(GOOSE) -dir migrations create $$name sql

migration-up:
	$(GOOSE) -dir migrations postgres "$(DATABASE_DSN)" up

migration-down:
	$(GOOSE) -dir migrations postgres "$(DATABASE_DSN)" down

migration-status:
	$(GOOSE) -dir migrations postgres "$(DATABASE_DSN)" status

generate:
	go generate ./...

lint:
	golangci-lint run
