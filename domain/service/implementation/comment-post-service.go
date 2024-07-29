package service

import (
	"strings"

	"github.com/google/uuid"

	commentDTO "github.com/cassiusbessa/vision-social-media/domain/service/dtos/comment"
	"github.com/cassiusbessa/vision-social-media/domain/service/errors"
	"github.com/cassiusbessa/vision-social-media/domain/service/mappers"
)

func (service *PostService) AddCommentToPost(command *commentDTO.AddCommentToPostCommand) (commentDTO.AddCommentToPostResponse, error) {

	uuidPost, err := uuid.Parse(command.PostID)
	if err != nil {
		return commentDTO.AddCommentToPostResponse{}, errors.NewInvalidArgument("Invalid post ID")
	}

	post, err := service.postRepo.GetPostByID(uuidPost)
	if err != nil {
		return commentDTO.AddCommentToPostResponse{}, err
	}
	if post == nil {
		return commentDTO.AddCommentToPostResponse{}, errors.NewResourceNotFound("Post not found")
	}

	comment, err := mappers.AddCommentToPostCommandToCommentEntity(*command)
	if err != nil {
		return commentDTO.AddCommentToPostResponse{}, err
	}

	comment.Validate()
	if len(comment.FailureMessage) > 0 {
		return commentDTO.AddCommentToPostResponse{}, errors.NewValidationError(strings.Join(comment.FailureMessage, ", "))
	}

	err = service.postRepo.AddCommentToPost(comment)
	if err != nil {
		return commentDTO.AddCommentToPostResponse{}, err
	}

	return commentDTO.AddCommentToPostResponse{
		CommentID: comment.ID.String(),
		Message:   "Comment saved",
	}, nil
}

func (service *PostService) RemovePostComment(command *commentDTO.RemovePostCommentCommand) (commentDTO.RemovePostCommentResponse, error) {
	uuidComment, err := uuid.Parse(command.CommentID)
	if err != nil {
		return commentDTO.RemovePostCommentResponse{}, errors.NewInvalidArgument("Invalid comment ID")
	}

	uuidPost, err := uuid.Parse(command.PostID)
	if err != nil {
		return commentDTO.RemovePostCommentResponse{}, errors.NewInvalidArgument("Invalid post ID")
	}

	removed, err := service.postRepo.RemovePostComment(uuidComment, uuidPost)
	if err != nil {
		return commentDTO.RemovePostCommentResponse{}, err
	}
	if !removed {
		return commentDTO.RemovePostCommentResponse{}, errors.NewResourceNotFound("Comment not found")
	}

	return commentDTO.RemovePostCommentResponse{
		Message: "Comment removed",
	}, nil
}

func (service *PostService) LoadPostCommentsByPostID(query *commentDTO.LoadOrderedCommentsQuery) ([]commentDTO.LoadedCommentResponse, error) {

	uuidPost, err := uuid.Parse(query.PostID)
	if err != nil {
		return []commentDTO.LoadedCommentResponse{}, errors.NewInvalidArgument("Invalid post ID")
	}

	comments, err := service.postRepo.LoadPostCommentsByPostID(uuidPost, query.Limit, query.Offset)
	if err != nil {
		return []commentDTO.LoadedCommentResponse{}, err
	}

	commentsDTO := make([]commentDTO.LoadedCommentResponse, len(comments))
	for i, comment := range comments {
		commentsDTO[i] = mappers.CommentEntityToLoadedCommentResponse(comment)
	}

	return commentsDTO, nil
}
