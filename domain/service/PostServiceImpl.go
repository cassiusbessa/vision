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
