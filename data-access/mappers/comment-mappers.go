package mappers

import (
	data "github.com/cassiusbessa/vision-social-media/data-access/sqlc-config"
	"github.com/cassiusbessa/vision-social-media/domain/core/entities"
	"github.com/google/uuid"
)

func LoadOrderedPostRowToPostComment(post data.LoadOrderedPostsRow) *entities.Comment {

	return entities.NewComment(
		entities.CommentWithID(post.CommentID.UUID),
		entities.CommentWithPostID(post.PostID),
		entities.CommentWithParentID(post.CommentParentID.UUID),
		entities.CommentWithUserID(post.CommentUserID.UUID),
		entities.CommentWithContent(post.CommentContent.String),
		entities.CommentWithReactions([]entities.Reaction{}),
		entities.CommentWithCreatedAt(post.CommentCreatedAt.Time),
		entities.CommentWithUpdatedAt(post.CommentUpdatedAt.Time),
	)
}

func CommentDbEntityToComment(comment data.Comment) *entities.Comment {

	return entities.NewComment(
		entities.CommentWithID(comment.ID),
		entities.CommentWithPostID(comment.PostID),
		entities.CommentWithParentID(comment.ParentID.UUID),
		entities.CommentWithUserID(comment.UserID),
		entities.CommentWithContent(comment.Content),
		entities.CommentWithReactions([]entities.Reaction{}),
		entities.CommentWithCreatedAt(comment.CreatedAt),
		entities.CommentWithUpdatedAt(comment.UpdatedAt),
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
		UserID:    comment.UserID,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt,
		UpdatedAt: comment.UpdatedAt,
	}
}
