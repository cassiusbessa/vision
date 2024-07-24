package data

import (
	"context"

	sqlc "github.com/cassiusbessa/vision-social-media/data-access/sqlc-config"
	"github.com/cassiusbessa/vision-social-media/domain/core/entities"
)

type PostRepository struct {
	queries *sqlc.Queries
}

func NewPostRepository(queries *sqlc.Queries) *PostRepository {
	return &PostRepository{
		queries: queries,
	}
}

func (repo *PostRepository) SavePost(post *entities.ProjectPost) error {
	_, err := repo.queries.CreatePost(context.Background(), projectEntityToQueryParams(post))
	return err
}
