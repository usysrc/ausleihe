# Ausleihe 
An app for loaning objects.

## How to start for development

Start the database:
```bash
docker compose up postgres -d
```

Migrate the database schema:
```bash
go run cmd/migrate/main.go up
```

Start the server:
```bash
./start.sh
```

This starts `air` that hot reloads golang builds.
