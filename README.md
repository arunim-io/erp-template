# ERP Template

## Prerequisites

- **Go 1.25+**
- **Docker** or **Podman** (for Postgres)
- **mise** (recommended) — installs Go, dbmate, sqlc, templ, air, golangci-lint, biome

Optional without mise: install [dbmate](https://github.com/amacneil/dbmate), [sqlc](https://sqlc.dev), [templ](https://templ.guide), [air](https://github.com/air-verse/air), and [golangci-lint](https://golangci-lint.run/) yourself.

## Setup

1. **Clone and enter the repo**

   ```bash
   cd erp-template
   ```

2. **Install tools (with mise)**

   ```bash
   mise install
   ```

3. **Environment**

   Copy the example env and adjust if needed:

   ```bash
   cp .env.example .env
   ```

   The app reads config from env vars with the `ERP_` prefix (see `.env.example`). For migrations, dbmate uses `DATABASE_URL` (same value as `ERP_DATABASE_URL` in the example). Mise loads `.env` automatically when you run tasks.

## Running the app

1. **Start the database** (Podman Compose via mise)

   ```bash
   mise run compose:up
   ```

2. **Run migrations**

   ```bash
   mise run db:migrate
   ```

3. **Generate code** (sqlc + templ)

   ```bash
   mise run generate
   ```

4. **Start the dev server**

   Hot reload (air); starts DB and waits for it, then runs the app:

   ```bash
   mise run dev
   ```

   The app listens on **http://localhost:8080** (air proxy); the app port is 8000.

## Mise tasks

| Task              | Command                 | Description                    |
|-------------------|-------------------------|--------------------------------|
| Start DB          | `mise run compose:up`   | Start Postgres (Podman)        |
| Stop DB           | `mise run compose:down` | Stop Postgres                  |
| Run migrations    | `mise run db:migrate`   | Apply dbmate migrations        |
| Generate code     | `mise run generate`    | sqlc generate + templ generate |
| Dev server        | `mise run dev`         | DB wait + air (hot reload)     |
| Tests             | `mise run test`        | `go test ./...`                |
| Lint              | `mise run lint`        | golangci-lint run              |
| Reset DB          | `mise run db:reset`    | Drop DB and re-run migrations  |
| Rollback migration| `mise run db:rollback` | Roll back last migration       |

## Testing

```bash
mise run test
```

## Linting

```bash
mise run lint
```

## Project layout

- `cmd/server/` — application entrypoint
- `internal/` — config, database, server, session, auth (stub), logging, templates
- `db/` — schema, migrations (dbmate), sqlc queries
- `static/` — CSS and other static assets (embedded in production)

## License

Use as you like.
