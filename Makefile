include .env
export

migration-create:
	@read -p "Enter migration name: " name; \
	goose -dir migrations create $$name sql

migration-up:
	goose -dir migrations postgres "$(DATABASE_DSN)" up

migration-down:
	goose -dir migrations postgres "$(DATABASE_DSN)" down

migration-status:
	goose -dir migrations postgres "$(DATABASE_DSN)" status

generate:
	go generate ./...

lint:
	golangci-lint run
