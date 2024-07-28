package service

import (
	"strings"

	dtos "github.com/cassiusbessa/vision-social-media/domain/service/dtos"
	errors "github.com/cassiusbessa/vision-social-media/domain/service/errors"
	"github.com/cassiusbessa/vision-social-media/domain/service/mappers"
	outputPorts "github.com/cassiusbessa/vision-social-media/domain/service/ports/output"
	"github.com/google/uuid"
)

type PostService struct {
	postRepo outputPorts.PostRepository
}

func NewPostService(postRepo outputPorts.PostRepository) *PostService {
	return &PostService{
		postRepo: postRepo,
	}
}

func (service *PostService) CreatePost(command *dtos.CreatePostCommand) (dtos.CreatedPostResponse, error) {

	post, err := mappers.CreatePostCommandToPostEntity(*command)
	if err != nil {
		return dtos.CreatedPostResponse{}, err
	}

	post.Validate()
	if len(post.FailureMessage) > 0 {
		return dtos.CreatedPostResponse{}, errors.NewValidationError(strings.Join(post.FailureMessage, ", "))
	}

	err = service.postRepo.SavePost(post)
	if err != nil {
		return dtos.CreatedPostResponse{}, err
	}

	return dtos.CreatedPostResponse{
		ID:      post.ID.String(),
		Message: "Post created",
	}, nil
}

func (service *PostService) UpdatePost(command *dtos.UpdatePostCommand) (dtos.UpdatedPostResponse, error) {

	uuidPost, err := uuid.Parse(command.ID)
	if err != nil {
		return dtos.UpdatedPostResponse{}, errors.NewInvalidArgument("Invalid post ID")
	}

	post, err := service.postRepo.GetPostByID(uuidPost)
	if err != nil {
		return dtos.UpdatedPostResponse{}, err
	}
	if post == nil {
		return dtos.UpdatedPostResponse{}, errors.NewResourceNotFound("Post not found")
	}

	updatedPost, err := mappers.UpdatePostCommandToPostEntity(*command, *post)
	if err != nil {
		return dtos.UpdatedPostResponse{}, err
	}

	post.Validate()
	if len(post.FailureMessage) > 0 {
		return dtos.UpdatedPostResponse{}, errors.NewValidationError(strings.Join(post.FailureMessage, ", "))
	}

	err = service.postRepo.UpdatePost(updatedPost)
	if err != nil {
		return dtos.UpdatedPostResponse{}, err
	}

	return dtos.UpdatedPostResponse{
		ID:      post.ID.String(),
		Message: "Post updated",
	}, nil
}

func (service *PostService) LoadOrderedPosts(query *dtos.LoadOrderedPostsQuery) ([]dtos.LoadedPostResponse, error) {

	posts, err := service.postRepo.LoadOrderedPosts()
	if err != nil {
		return nil, err
	}

	loadedPosts := make([]dtos.LoadedPostResponse, 0)
	for _, post := range posts {
		loadedPosts = append(loadedPosts, mappers.PostEntityToLoadedPostResponse(post))
	}

	return loadedPosts, nil
}

func (service *PostService) ReactToPost(command *dtos.ReactToPostCommand) (dtos.ReactToPostResponse, error) {

	uuidPost, err := uuid.Parse(command.PostID)
	if err != nil {
		return dtos.ReactToPostResponse{}, errors.NewInvalidArgument("Invalid post ID")
	}

	post, err := service.postRepo.GetPostByID(uuidPost)
	if err != nil {
		return dtos.ReactToPostResponse{}, err
	}
	if post == nil {
		return dtos.ReactToPostResponse{}, errors.NewResourceNotFound("Post not found")
	}
	for _, reaction := range post.Reactions {
		if reaction.UserID == uuid.MustParse(command.UserID) {
			return dtos.ReactToPostResponse{}, errors.NewResourceAlreadyExists("User already reacted to this post")
		}
	}

	reaction, err := mappers.ReactToPostCommandToReactionEntity(*command)
	if err != nil {
		return dtos.ReactToPostResponse{}, err
	}

	reaction.Validate()
	if len(reaction.FailureMessage) > 0 {
		return dtos.ReactToPostResponse{}, errors.NewValidationError(strings.Join(reaction.FailureMessage, ", "))
	}

	err = service.postRepo.AddReactionToPost(reaction)
	if err != nil {
		return dtos.ReactToPostResponse{}, err
	}

	return dtos.ReactToPostResponse{
		ID:      reaction.ID.String(),
		Message: "Reaction saved",
	}, nil
}

func (service *PostService) RemovePostReaction(command *dtos.RemovePostReactionCommand) (dtos.RemovePostReactionResponse, error) {
	uuidReaction, err := uuid.Parse(command.ReactionID)
	if err != nil {
		return dtos.RemovePostReactionResponse{}, errors.NewInvalidArgument("Invalid reaction ID")
	}

	uuidPost, err := uuid.Parse(command.PostID)
	if err != nil {
		return dtos.RemovePostReactionResponse{}, errors.NewInvalidArgument("Invalid post ID")
	}

	removed, err := service.postRepo.RemovePostReaction(uuidReaction, uuidPost)
	if err != nil {
		return dtos.RemovePostReactionResponse{}, err
	}
	if !removed {
		return dtos.RemovePostReactionResponse{}, errors.NewResourceNotFound("Reaction not found")
	}

	return dtos.RemovePostReactionResponse{
		Message: "Reaction removed",
	}, nil
}

func (service *PostService) AddCommentToPost(command *dtos.AddCommentToPostCommand) (dtos.AddCommentToPostResponse, error) {

	uuidPost, err := uuid.Parse(command.PostID)
	if err != nil {
		return dtos.AddCommentToPostResponse{}, errors.NewInvalidArgument("Invalid post ID")
	}

	post, err := service.postRepo.GetPostByID(uuidPost)
	if err != nil {
		return dtos.AddCommentToPostResponse{}, err
	}
	if post == nil {
		return dtos.AddCommentToPostResponse{}, errors.NewResourceNotFound("Post not found")
	}

	comment, err := mappers.AddCommentToPostCommandToCommentEntity(*command)
	if err != nil {
		return dtos.AddCommentToPostResponse{}, err
	}

	comment.Validate()
	if len(comment.FailureMessage) > 0 {
		return dtos.AddCommentToPostResponse{}, errors.NewValidationError(strings.Join(comment.FailureMessage, ", "))
	}

	err = service.postRepo.AddCommentToPost(comment)
	if err != nil {
		return dtos.AddCommentToPostResponse{}, err
	}

	return dtos.AddCommentToPostResponse{
		CommentID: comment.ID.String(),
		Message:   "Comment saved",
	}, nil
}
