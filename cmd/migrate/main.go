package main

import (
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	m, err := migrate.New(
		"file://cmd/migrate/migrations",                                   // source URL
		"postgresql://username:password@localhost/dbname?sslmode=disable", // database URL
	)
	if err != nil {
		log.Fatal(err.Error())
	}

	cmd := os.Args[len(os.Args)-1]

	if cmd == "up" {
		err = m.Up()
		if err != nil {
			log.Fatal(err.Error())
		}
	}

	if cmd == "down" {
		err = m.Down()
		if err != nil {
			log.Fatal(err.Error())
		}
	}
}
