package service

import (
	"strings"

	reactionDTO "github.com/cassiusbessa/vision-social-media/domain/service/dtos/reaction"
	"github.com/cassiusbessa/vision-social-media/domain/service/errors"
	"github.com/cassiusbessa/vision-social-media/domain/service/mappers"
	"github.com/google/uuid"
)

func (service *PostService) ReactToPost(command *reactionDTO.ReactToPostCommand) (reactionDTO.ReactToPostResponse, error) {

	uuidPost, err := uuid.Parse(command.PostID)
	if err != nil {
		return reactionDTO.ReactToPostResponse{}, errors.NewInvalidArgument("Invalid post ID")
	}

	post, err := service.postRepo.GetPostByID(uuidPost)
	if err != nil {
		return reactionDTO.ReactToPostResponse{}, err
	}
	if post == nil {
		return reactionDTO.ReactToPostResponse{}, errors.NewResourceNotFound("Post not found")
	}
	for _, reaction := range post.Reactions {
		if reaction.Author.ID == uuid.MustParse(command.UserID) {
			return reactionDTO.ReactToPostResponse{}, errors.NewResourceAlreadyExists("User already reacted to this post")
		}
	}

	reaction, err := mappers.ReactToPostCommandToReactionEntity(*command)
	if err != nil {
		return reactionDTO.ReactToPostResponse{}, err
	}

	reaction.Validate()
	if len(reaction.FailureMessage) > 0 {
		return reactionDTO.ReactToPostResponse{}, errors.NewValidationError(strings.Join(reaction.FailureMessage, ", "))
	}

	err = service.postRepo.AddReactionToPost(reaction)
	if err != nil {
		return reactionDTO.ReactToPostResponse{}, err
	}

	return reactionDTO.ReactToPostResponse{
		ID:      reaction.ID.String(),
		Message: "Reaction saved",
	}, nil
}

func (service *PostService) RemovePostReaction(command *reactionDTO.RemovePostReactionCommand) (reactionDTO.RemovePostReactionResponse, error) {
	uuidReaction, err := uuid.Parse(command.ReactionID)
	if err != nil {
		return reactionDTO.RemovePostReactionResponse{}, errors.NewInvalidArgument("Invalid reaction ID")
	}

	uuidPost, err := uuid.Parse(command.PostID)
	if err != nil {
		return reactionDTO.RemovePostReactionResponse{}, errors.NewInvalidArgument("Invalid post ID")
	}

	removed, err := service.postRepo.RemovePostReaction(uuidReaction, uuidPost)
	if err != nil {
		return reactionDTO.RemovePostReactionResponse{}, err
	}
	if !removed {
		return reactionDTO.RemovePostReactionResponse{}, errors.NewResourceNotFound("Reaction not found")
	}

	return reactionDTO.RemovePostReactionResponse{
		Message: "Reaction removed",
	}, nil
}

func (service *PostService) LoadPostReactionsByPostID(query *reactionDTO.LoadOrderedReactionsQuery) ([]reactionDTO.LoadReactionResponse, error) {

	uuidPost, err := uuid.Parse(query.PostID)
	if err != nil {
		return []reactionDTO.LoadReactionResponse{}, errors.NewInvalidArgument("Invalid post ID")
	}

	reactions, err := service.postRepo.LoadReactionsByPostID(uuidPost, query.Limit, query.Offset)
	if err != nil {
		return []reactionDTO.LoadReactionResponse{}, err
	}

	loadedReactions := make([]reactionDTO.LoadReactionResponse, 0)
	for _, reaction := range reactions {
		loadedReactions = append(loadedReactions, mappers.ReactionEntityToLoadedReactionResponse(reaction))
	}

	return loadedReactions, nil
}
