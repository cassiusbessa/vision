package mappers

import (
	data "github.com/cassiusbessa/vision-social-media/data-access/sqlc-config"
	"github.com/cassiusbessa/vision-social-media/domain/core/entities"
	"github.com/google/uuid"
)

func CommentDbEntityToComment(comment data.Comment) *entities.Comment {

	return entities.NewComment(
		entities.CommentWithID(comment.ID),
		entities.CommentWithPostID(comment.PostID),
		entities.CommentWithParentID(comment.ParentID.UUID),
		entities.CommentWithAuthorID(comment.UserID),
		entities.CommentWithContent(comment.Content),
		entities.CommentWithReactions([]entities.Reaction{}),
		entities.CommentWithCreatedAt(comment.CreatedAt),
		entities.CommentWithUpdatedAt(comment.UpdatedAt),
	)
}

func LoadedCommentToComment(comment data.LoadCommentsByPostIDRow) *entities.Comment {

	var parentID uuid.UUID

	if comment.ParentID.Valid {
		parentID = comment.ParentID.UUID
	}

	return entities.NewComment(
		entities.CommentWithID(comment.CommentID),
		entities.CommentWithAuthorID(comment.UserID),
		entities.CommentWithAuthorImage(comment.AuthorImage.String),
		entities.CommentWithAuthorName(comment.AuthorName),
		entities.CommentWithPostID(comment.PostID),
		entities.CommentWithParentID(parentID),
		entities.CommentWithAuthorID(comment.UserID),
		entities.CommentWithContent(comment.CommentContent),
		entities.CommentWithCreatedAt(comment.CommentCreatedAt),
		entities.CommentWithUpdatedAt(comment.CommentUpdatedAt),
	)
}

func CommentEntityToCreateQueryParams(comment *entities.Comment) data.CreateCommentParams {

	var parentID uuid.NullUUID

	if comment.ParentID != uuid.Nil {
		parentID = uuid.NullUUID{
			UUID:  comment.ParentID,
			Valid: true,
		}
	}
	return data.CreateCommentParams{
		ID:        comment.ID,
		PostID:    comment.PostID,
		ParentID:  parentID,
		UserID:    comment.Author.ID,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
	}
}
