package ports

import (
	entities "github.com/cassiusbessa/vision-social-media/domain/core/entities"
	"github.com/google/uuid"
)

type PostRepository interface {
	SavePost(post *entities.ProjectPost) error
	UpdatePostByProjectID(post *entities.ProjectPost) error
	RemovePostByProjectID(projectID uuid.UUID) (bool, error)
	GetPostByID(postID uuid.UUID) (*entities.ProjectPost, error)
	GetPostByProjectID(projectID uuid.UUID) (*entities.ProjectPost, error)
	LoadOrderedPosts(limit, offSet int32) ([]entities.ProjectPost, error)
	AddReactionToPost(reaction *entities.Reaction) error
	RemovePostReaction(reactionID, postID uuid.UUID) (bool, error)
	LoadReactionsByPostID(postID uuid.UUID, limit, offSet int32) ([]entities.Reaction, error)
	AddCommentToPost(comment *entities.Comment) error
	RemovePostComment(commentID, postID uuid.UUID) (bool, error)
	LoadPostCommentsByPostID(postID uuid.UUID, limit, offSet int32) ([]entities.Comment, error)
}
