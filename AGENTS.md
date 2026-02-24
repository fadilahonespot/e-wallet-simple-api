# AGENTS.md

## Cursor Cloud specific instructions

This is a Go (Gin) REST API for an e-wallet system backed by PostgreSQL. Architecture: Handler → Usecase → Repo → Database (clean architecture).

### Services

| Service | How to run |
|---------|-----------|
| PostgreSQL | `docker run -d --name ewallet-postgres -e POSTGRES_DB=ewallet -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=pas12345 -p 5432:5432 postgres:alpine` |
| Go API server | Export env vars then `go run main.go` (see below) |

### Running the server locally

The `.env` file has different credentials than what's needed for Docker-based PostgreSQL. When running locally (not via `docker-compose`), you must export environment variables **before** starting the server so that `godotenv/autoload` does not overwrite them with stale `.env` values:

```sh
export POSTGRES_HOST=localhost POSTGRES_DB=ewallet POSTGRES_USER=postgres POSTGRES_PASSWORD=pas12345 PORT=8090
go run main.go
```

### Key commands

See `Makefile` for all targets. Most useful:

- **Unit tests** (no DB needed): `make unit-test` — uses sqlmock, runs with `-short` flag
- **Build**: `go build -o main .`
- **Dependencies**: `go get -v ./...`

### Gotchas

- `database.sql` is loaded at startup via `setting.ConnectDB()`. It uses bare `CREATE` (not `CREATE OR REPLACE` / `IF NOT EXISTS` for tables/functions), so restarting against an existing database will print SQL errors to stdout — these are non-fatal and can be ignored.
- No linter is configured in the project. Use `go vet ./...` for basic static analysis.
- The project has no pre-commit hooks or CI configuration.
