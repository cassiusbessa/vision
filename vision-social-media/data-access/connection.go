package data

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
)

func NewDbConn() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), "postgres://root:root@db:5432/vision?sslmode=disable")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	return conn
}
