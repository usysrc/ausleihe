#!/bin/sh

# start the database
docker compose up postgres -d --wait

# stop the database when this script exits
trap 'docker compose down' EXIT

# migrate the database
go run cmd/migrate/main.go

# start the server
air --build.cmd "go build -o ./bin/api ./cmd/app/" --build.bin "./bin/api"
