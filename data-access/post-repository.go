package data

import (
	"context"

	sqlc "github.com/cassiusbessa/vision-social-media/data-access/sqlc-config"
	"github.com/jackc/pgx/v4"
)

type PostRepository struct {
	queries *sqlc.Queries
	db      *pgx.Conn
}

func NewPostRepository(queries *sqlc.Queries, db *pgx.Conn) *PostRepository {
	return &PostRepository{
		queries: queries,
		db:      db,
	}
}

func withTransaction(ctx context.Context, db *pgx.Conn, fn func(context.Context, *sqlc.Queries) error) error {
	tx, err := db.Begin(ctx)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback(ctx)
			panic(p)
		} else if err != nil {
			tx.Rollback(ctx)
		}
	}()

	qtx := sqlc.New(tx)

	err = fn(ctx, qtx)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}
