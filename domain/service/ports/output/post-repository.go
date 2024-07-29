package ports

import (
	entities "github.com/cassiusbessa/vision-social-media/domain/core/entities"
	"github.com/google/uuid"
)

type PostRepository interface {
	SavePost(post *entities.ProjectPost) error
	UpdatePost(post *entities.ProjectPost) error
	RemovePost(postID uuid.UUID) (bool, error)
	GetPostByID(postID uuid.UUID) (*entities.ProjectPost, error)
	LoadOrderedPosts(limit, offSet int32) ([]entities.ProjectPost, error)
	AddReactionToPost(reaction *entities.Reaction) error
	RemovePostReaction(reactionID, postID uuid.UUID) (bool, error)
	AddCommentToPost(comment *entities.Comment) error
	RemovePostComment(commentID, postID uuid.UUID) (bool, error)
}
