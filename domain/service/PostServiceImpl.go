package service

import (
	"strings"

	entities "github.com/cassiusbessa/vision-social-media/domain/core/entities"
	dtos "github.com/cassiusbessa/vision-social-media/domain/service/dtos"
	errors "github.com/cassiusbessa/vision-social-media/domain/service/errors"
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

	uuidProject, err := uuid.Parse(command.ProjectID)
	if err != nil {
		return dtos.CreatedPostResponse{}, errors.NewInvalidArgument("Invalid project ID")
	}

	uuidUser, err := uuid.Parse(command.AuthorID)
	if err != nil {
		return dtos.CreatedPostResponse{}, errors.NewInvalidArgument("Invalid user ID")
	}
	post := entities.NewProjectPost(
		uuidProject,
		uuidUser,
		command.Title,
		command.Content,
		command.RepoLink,
		command.DemoLink,
		command.PostImage,
	)
	post.Validate()
	if len(post.FailureMessage) > 0 {
		return dtos.CreatedPostResponse{}, errors.NewValidationError(strings.Join(post.FailureMessage, ", "))
	}

	err = service.postRepo.SavePost(post)
	if err != nil {
		return dtos.CreatedPostResponse{}, err
	}

	return dtos.CreatedPostResponse{}, nil
}
