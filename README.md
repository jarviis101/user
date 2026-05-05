# User Microservice API

A Go-based microservice for user management with dual transport support: GraphQL (HTTP) and gRPC, backed by PostgreSQL.

## Stack

| Name       | Version   | Description                        |
|------------|-----------|------------------------------------|
| Go         | 1.26.2    | Primary language                   |
| PostgreSQL | 16        | Relational database                |
| Echo       | v4        | HTTP server framework              |
| GraphQL    | gqlgen    | GraphQL API layer                  |
| gRPC       | v1.81.0   | High-performance RPC framework     |
| pgx        | v5.9.2    | Native PostgreSQL driver           |
| Goose      | v3.27.1   | Database migration tool            |

## Requirements

- [Go 1.25+](https://go.dev/)
- [Docker & Docker Compose](https://docs.docker.com/compose/)
- [Make](https://www.gnu.org/software/make/)
- [Goose](https://github.com/pressly/goose) — `go install github.com/pressly/goose/v3/cmd/goose@latest`

## Environment variables

| Variable       | Default                                                       | Description                      |
|----------------|---------------------------------------------------------------|----------------------------------|
| `APP_KEY`      | —                                                             | API authentication key (required)|
| `APP_HOST`     | `localhost`                                                   | Server host                      |
| `APP_PORT`     | `8000`                                                        | Server port                      |
| `APP_DEBUG`    | `1`                                                           | Enable GraphQL Playground        |
| `DATABASE_DSN` | `postgres://user:password@localhost:5432/user?sslmode=disable`| PostgreSQL connection string     |

## Local setup

```bash
cp .env.example .env
# fill in the required environment variables

docker compose up -d    # start PostgreSQL
make migration-up       # apply database migrations

go run cmd/http/main.go # start HTTP server (GraphQL)
# or
go run cmd/grpc/main.go # start gRPC server
```

## Make targets

| Command               | Description                        |
|-----------------------|------------------------------------|
| `make migration-create` | Create a new migration file      |
| `make migration-up`     | Apply pending migrations          |
| `make migration-down`   | Roll back the last migration      |
| `make migration-status` | Show migration status             |
| `make generate`         | Regenerate GraphQL & protobuf code|
| `make lint`             | Run golangci-lint                 |
