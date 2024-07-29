package data

import (
	"context"

	"github.com/cassiusbessa/vision-social-media/data-access/mappers"
	sqlc "github.com/cassiusbessa/vision-social-media/data-access/sqlc-config"
	"github.com/cassiusbessa/vision-social-media/domain/core/entities"
	"github.com/google/uuid"
)

func (repo *PostRepository) SavePost(post *entities.ProjectPost) error {
	err := repo.queries.CreatePost(context.Background(), mappers.PostEntityToCreateQueryParams(post))
	return err
}

func (repo *PostRepository) RemovePost(postID uuid.UUID) (bool, error) {
	err := repo.queries.DeletePostById(context.Background(), postID)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (repo *PostRepository) UpdatePost(post *entities.ProjectPost) error {
	err := repo.queries.UpdatePost(context.Background(), mappers.PostEntityToUpdateQueryParams(post))
	return err
}

func (repo *PostRepository) GetPostByID(postID uuid.UUID) (*entities.ProjectPost, error) {

	post, err := repo.queries.GetPostByID(context.Background(), postID)
	if err != nil {
		return nil, err
	}

	comments, err := repo.queries.GetCommentsByPostID(context.Background(), postID)
	if err != nil {
		return nil, err
	}

	reactions, err := repo.queries.GetReactionsByPostID(context.Background(), postID)
	if err != nil {
		return nil, err
	}

	return mappers.PostDBEntityToProjectPost(post, comments, reactions), nil
}

func (repo *PostRepository) LoadOrderedPosts(limit, offSet int32) ([]entities.ProjectPost, error) {
	dbPosts, err := repo.queries.LoadOrderedPosts(context.Background(), sqlc.LoadOrderedPostsParams{
		Limit:  limit,
		Offset: offSet,
	})
	if err != nil {
		return nil, err
	}

	var posts []entities.ProjectPost
	for _, dbPost := range dbPosts {
		posts = append(posts, *mappers.LoadOrderedPostRowToProjectPosts(dbPost))
	}

	return posts, nil

}
