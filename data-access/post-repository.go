package data

import (
	"context"

	sqlc "github.com/cassiusbessa/vision-social-media/data-access/sqlc-config"
	"github.com/cassiusbessa/vision-social-media/domain/core/entities"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
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
	_, err := repo.queries.CreatePost(context.Background(), projectEntityToCreateQueryParams(post))
	return err
}

func (repo *PostRepository) UpdatePost(post *entities.ProjectPost) error {
	_, err := repo.queries.UpdatePost(context.Background(), projectEntityToUpdateQueryParams(post))
	return err
}

func (repo *PostRepository) GetPostByID(postID uuid.UUID) (*entities.ProjectPost, error) {
	var id pgtype.UUID
	id.Bytes = [16]byte(postID)
	id.Valid = true

	post, err := repo.queries.GetPostByID(context.Background(), id)
	if err != nil {
		return nil, err
	}

	return postDBEntityToProjectPost(post), nil
}
