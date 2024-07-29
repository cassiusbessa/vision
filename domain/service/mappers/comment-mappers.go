package mappers

import (
	"time"

	"github.com/cassiusbessa/vision-social-media/domain/core/entities"
	commentDTO "github.com/cassiusbessa/vision-social-media/domain/service/dtos/comment"
	"github.com/cassiusbessa/vision-social-media/domain/service/errors"
	"github.com/google/uuid"
)

func commentEntityToLoadedCommentResponse(comment entities.Comment) commentDTO.LoadCommentResponse {

	var parentID string
	if comment.ParentID == uuid.Nil {
		parentID = ""
	} else {
		parentID = comment.ParentID.String()
	}

	return commentDTO.LoadCommentResponse{
		ID:        comment.ID.String(),
		AuthorID:  comment.UserID.String(),
		PostID:    comment.PostID.String(),
		ParentID:  parentID,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt.Format(time.RFC3339),
		UpdatedAt: comment.UpdatedAt.Format(time.RFC3339),
	}
}

func AddCommentToPostCommandToCommentEntity(command commentDTO.AddCommentToPostCommand) (*entities.Comment, error) {
	uuidPost, err := uuid.Parse(command.PostID)
	if err != nil {
		return &entities.Comment{}, errors.NewInvalidArgument("Invalid post ID")
	}

	uuidUser, err := uuid.Parse(command.AuthorID)
	if err != nil {
		return &entities.Comment{}, errors.NewInvalidArgument("Invalid user ID")
	}

	comment := entities.NewComment(
		entities.CommentWithID(uuid.New()),
		entities.CommentWithPostID(uuidPost),
		entities.CommentWithUserID(uuidUser),
		entities.CommentWithContent(command.Content),
		entities.CommentWithCreatedAt(time.Now()),
		entities.CommentWithUpdatedAt(time.Now()),
	)

	return comment, nil
}
