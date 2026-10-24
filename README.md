# Todo List — Go Web Application

A todo-list application built with Go, Gin, GORM, and PostgreSQL. It includes a server-rendered HTML interface with JavaScript and CSS assets.

## Structure

- `router/`, `handler/`: HTTP routing and handlers.
- `services/`, `Repositories/`: application and persistence logic.
- `entitys/`, `dto/`: database entities and request/response types.
- `db/`: database setup and automatic schema migration.
- `templates/`, `static/`: browser interface.

## Run locally

Requirements: Go 1.22.5 or newer and PostgreSQL.

Create a local database named `todo`, then set the PostgreSQL environment variables for your instance:

```powershell
$env:PGHOST = 'localhost'
$env:PGPORT = '5432'
$env:PGDATABASE = 'todo'
$env:PGUSER = 'postgres'
$env:PGPASSWORD = '<your-local-password>'

go mod download
go run .
```

Check the server output for its listening address. The application calls GORM's `AutoMigrate` on startup, so the configured database user needs permission to create or update the todo table.

Credentials belong in your local environment. The database connection currently disables TLS for local development; configure TLS before using a remote database.

`.env.example` documents the variable names. The application does not automatically load `.env` files; export the variables in your shell before starting it.
