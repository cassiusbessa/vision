package service

import (
	"strings"

	postDTO "github.com/cassiusbessa/vision-social-media/domain/service/dtos/post"
	"github.com/cassiusbessa/vision-social-media/domain/service/errors"
	"github.com/cassiusbessa/vision-social-media/domain/service/mappers"
	"github.com/google/uuid"
)

func (service *PostService) CreatePost(command *postDTO.CreatePostCommand) (postDTO.CreatedPostResponse, error) {

	post, err := mappers.CreatePostCommandToPostEntity(*command)
	if err != nil {
		return postDTO.CreatedPostResponse{}, err
	}

	post.Validate()
	if len(post.FailureMessage) > 0 {
		return postDTO.CreatedPostResponse{}, errors.NewValidationError(strings.Join(post.FailureMessage, ", "))
	}

	err = service.postRepo.SavePost(post)
	if err != nil {
		return postDTO.CreatedPostResponse{}, err
	}

	return postDTO.CreatedPostResponse{
		ID:      post.ID.String(),
		Message: "Post created",
	}, nil
}

func (service *PostService) DeletePost(command *postDTO.RemovePostCommand) (postDTO.RemovedPostResponse, error) {

	uuidProject, err := uuid.Parse(command.ProjectID)
	if err != nil {
		return postDTO.RemovedPostResponse{}, errors.NewInvalidArgument("Invalid post ID")
	}

	removed, err := service.postRepo.RemovePostByProjectID(uuidProject)
	if err != nil {
		return postDTO.RemovedPostResponse{}, err
	}
	if !removed {
		return postDTO.RemovedPostResponse{}, errors.NewResourceNotFound("Post not found")
	}

	return postDTO.RemovedPostResponse{
		Message: "Post removed",
	}, nil
}

func (service *PostService) UpdatePost(command *postDTO.UpdatePostCommand) (postDTO.UpdatedPostResponse, error) {

	projectID, err := uuid.Parse(command.ProjectID)
	if err != nil {
		return postDTO.UpdatedPostResponse{}, errors.NewInvalidArgument("Invalid project ID")
	}

	post, err := service.postRepo.GetPostByProjectID(projectID)
	if err != nil {
		return postDTO.UpdatedPostResponse{}, err
	}
	if post == nil {
		return postDTO.UpdatedPostResponse{}, errors.NewResourceNotFound("Post not found")
	}

	updatedPost, err := mappers.UpdatePostCommandToPostEntity(*command, *post)
	if err != nil {
		return postDTO.UpdatedPostResponse{}, err
	}

	post.Validate()
	if len(post.FailureMessage) > 0 {
		return postDTO.UpdatedPostResponse{}, errors.NewValidationError(strings.Join(post.FailureMessage, ", "))
	}

	err = service.postRepo.UpdatePostByProjectID(updatedPost)
	if err != nil {
		return postDTO.UpdatedPostResponse{}, err
	}

	return postDTO.UpdatedPostResponse{
		ID:      post.ID.String(),
		Message: "Post updated",
	}, nil
}

func (service *PostService) LoadOrderedPosts(query *postDTO.LoadOrderedPostsQuery) ([]postDTO.LoadedPostResponse, error) {

	posts, err := service.postRepo.LoadOrderedPosts(query.Limit, query.Offset)
	if err != nil {
		return nil, err
	}

	loadedPosts := make([]postDTO.LoadedPostResponse, 0)
	for _, post := range posts {
		loadedPosts = append(loadedPosts, mappers.PostEntityToLoadedPostResponse(post))
	}

	return loadedPosts, nil
}
