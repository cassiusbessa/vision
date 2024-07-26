package data

import (
	"context"

	sqlc "github.com/cassiusbessa/vision-social-media/data-access/sqlc-config"
	"github.com/cassiusbessa/vision-social-media/domain/core/entities"
	"github.com/google/uuid"
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

	post, err := repo.queries.GetPostByID(context.Background(), postID)
	if err != nil {
		return nil, err
	}

	return postDBEntityToProjectPost(post), nil
}

func (repo *PostRepository) LoadOrderedPosts() ([]entities.ProjectPost, error) {
	posts, err := repo.queries.LoadOrderedPosts(context.Background())
	if err != nil {
		return nil, err
	}

	entitiesPosts := map[uuid.UUID]*entities.ProjectPost{}

	for _, post := range posts {
		if _, ok := entitiesPosts[post.PostID]; !ok {
			entitiesPosts[post.PostID] = loadOrderedPostRowToProjectPosts(post)
		}

		if post.CommentID.Valid {
			entitiesPosts[post.PostID].AddComment(loadOrderedPostRowToProjectComment(post))
		}

		if post.ReactionID.Valid {
			entitiesPosts[post.PostID].AddReaction(loadOrderedPostRowToProjectReaction(post))
		}
	}

	result := make([]entities.ProjectPost, 0, len(entitiesPosts))
	for _, post := range entitiesPosts {
		result = append(result, *post)
	}

	return result, nil
}
