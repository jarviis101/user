# User Microservice API

A Go-based microservice for user management, exposing a GraphQL API backed by PostgreSQL.

## Stack

| Name       | Version  | Description                        |
|------------|----------|------------------------------------|
| Go         | 1.26.2   | Primary language                   |
| PostgreSQL | 16       | Relational database                |
| GraphQL    | gqlgen   | API layer (queries & mutations)    |
| Echo       | v4       | HTTP server framework              |
| Goose      | 3.27.1   | Database migration tool            |

## Requirements

- [Go 1.25+](https://go.dev/)
- [Docker & Docker Compose](https://docs.docker.com/compose/)
- [Make](https://www.gnu.org/software/make/)
- [Goose](https://github.com/pressly/goose) — `go install github.com/pressly/goose/v3/cmd/goose@latest`

## Environment variables

| Variable       | Description                          |
|----------------|--------------------------------------|
| `APP_KEY`      | Application secret key               |
| `APP_HOST`     | Server host (e.g. `0.0.0.0`)        |
| `APP_PORT`     | Server port (e.g. `8080`)           |
| `APP_DEBUG`    | Debug mode (`true` / `false`)        |
| `DATABASE_DSN` | PostgreSQL DSN connection string     |

## Local deployment

```bash
cp .env.example .env
# fill in the required environment variables

docker compose up -d       
make migration-up           
go run cmd/app/main.go     
```
