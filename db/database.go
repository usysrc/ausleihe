package db

import (
	"context"

	"github.com/jackc/pgx/v4"
)

var Postgres *pgx.Conn

func CreateDatabase() (*pgx.Conn, error) {
	db, err := pgx.Connect(context.Background(), "postgresql://username:password@localhost/dbname")
	if err != nil {
		return nil, err
	}
	Postgres = db
	return db, nil
}

func CloseDatabase(db *pgx.Conn) {
	defer db.Close(context.Background())
}
