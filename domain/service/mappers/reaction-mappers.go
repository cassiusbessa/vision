package mappers

import (
	"time"

	"github.com/cassiusbessa/vision-social-media/domain/core/entities"
	reactionDTO "github.com/cassiusbessa/vision-social-media/domain/service/dtos/reaction"
	"github.com/cassiusbessa/vision-social-media/domain/service/errors"
	"github.com/google/uuid"
)

func ReactionEntityToLoadedReactionResponse(reaction entities.Reaction) reactionDTO.LoadReactionResponse {

	var parentID string
	if reaction.ParentID.UUID == uuid.Nil {
		parentID = ""
	} else {
		parentID = reaction.ParentID.UUID.String()
	}

	return reactionDTO.LoadReactionResponse{
		ID:        reaction.ID.String(),
		UserID:    reaction.UserID.String(),
		PostID:    reaction.PostID.String(),
		ParentID:  parentID,
		Type:      reaction.Type,
		CreatedAt: reaction.CreatedAt.Format(time.RFC3339),
	}
}

func ReactToPostCommandToReactionEntity(command reactionDTO.ReactToPostCommand) (*entities.Reaction, error) {
	uuidPost, err := uuid.Parse(command.PostID)
	if err != nil {
		return &entities.Reaction{}, errors.NewInvalidArgument("Invalid post ID")
	}

	uuidUser, err := uuid.Parse(command.UserID)
	if err != nil {
		return &entities.Reaction{}, errors.NewInvalidArgument("Invalid user ID")
	}

	var reactionType entities.ReactionType

	switch command.Type {
	case "Like":
		reactionType = entities.Like
	case "Dislike":
		reactionType = entities.Dislike
	case "Love":
		reactionType = entities.Love
	case "Wow":
		reactionType = entities.Wow
	case "Angry":
		reactionType = entities.Angry
	default:
		return &entities.Reaction{}, errors.NewInvalidArgument("Invalid reaction type")
	}

	parentId := uuid.NullUUID{
		UUID:  uuid.Nil,
		Valid: false,
	}

	reaction := entities.NewReaction(
		entities.ReactionWithID(uuid.New()),
		entities.ReactionWithPostID(uuidPost),
		entities.ReactionWithUserID(uuidUser),
		entities.ReactionWithReactionType(reactionType),
		entities.ReactionWithParentID(parentId),
		entities.ReactionWithCreatedAt(time.Now()),
	)

	return reaction, nil
}
