package data

import (
	"context"

	"github.com/cassiusbessa/vision-social-media/data-access/mappers"
	sqlc "github.com/cassiusbessa/vision-social-media/data-access/sqlc-config"
	"github.com/cassiusbessa/vision-social-media/domain/core/entities"
	"github.com/google/uuid"
)

func (repo *PostRepository) AddCommentToPost(comment *entities.Comment) error {
	return withTransaction(context.Background(), repo.db, func(ctx context.Context, qtx *sqlc.Queries) error {
		err := qtx.CreateComment(ctx, mappers.CommentEntityToCreateQueryParams(comment))
		if err != nil {
			return err
		}

		err = qtx.AddCommentCount(ctx, comment.PostID)
		if err != nil {
			return err
		}

		return nil
	})
}

func (repo *PostRepository) RemovePostComment(commentID, postID uuid.UUID) (bool, error) {
	return true, withTransaction(context.Background(), repo.db, func(ctx context.Context, qtx *sqlc.Queries) error {
		err := qtx.DeleteCommentById(ctx, commentID)
		if err != nil {
			return err
		}

		err = qtx.RemoveCommentCount(ctx, postID)
		if err != nil {
			return err
		}

		return nil
	})
}

func (repo *PostRepository) LoadPostCommentsByPostID(postID uuid.UUID, limit, offSet int32) ([]entities.Comment, error) {
	comments, err := repo.queries.LoadCommentsByPostID(context.Background(), sqlc.LoadCommentsByPostIDParams{
		PostID: postID,
		Limit:  limit,
		Offset: offSet,
	})
	if err != nil {
		return nil, err
	}

	var commentsEntities []entities.Comment

	for _, comment := range comments {
		commentsEntities = append(commentsEntities, *mappers.LoadedCommentToComment(comment))
	}

	return commentsEntities, nil
}
