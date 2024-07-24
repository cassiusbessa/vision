package data

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func NewDbConn() *pgx.Conn {
	conn, err := pgx.Connect(context.Background(), "postgres://localhost:5432/vision?sslmode=disable")
	if err != nil {
		log.Fatalf("Unable to connect to database: %v\n", err)
	}
	return conn
}
