package data

import (
	"context"

	"github.com/cassiusbessa/vision-social-media/data-access/mappers"
	sqlc "github.com/cassiusbessa/vision-social-media/data-access/sqlc-config"
	"github.com/cassiusbessa/vision-social-media/domain/core/entities"
	"github.com/google/uuid"
)

func (repo *PostRepository) AddReactionToPost(reaction *entities.Reaction) error {

	return withTransaction(context.Background(), repo.db, func(ctx context.Context, qtx *sqlc.Queries) error {
		err := qtx.CreateReaction(ctx, mappers.ReactionEntityToCreateReactionQueryParams(reaction))
		if err != nil {
			return err
		}

		err = qtx.AddReactionCount(ctx, reaction.PostID)
		if err != nil {
			return err
		}

		return nil
	})
}

func (repo *PostRepository) RemovePostReaction(reactionID, postID uuid.UUID) (bool, error) {
	return true, withTransaction(context.Background(), repo.db, func(ctx context.Context, qtx *sqlc.Queries) error {
		err := qtx.DeleteReactionById(ctx, reactionID)
		if err != nil {
			return err
		}

		err = qtx.RemoveReactionCount(ctx, postID)
		if err != nil {
			return err
		}

		return nil
	})
}
