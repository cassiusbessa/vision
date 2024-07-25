package ports

import (
	entities "github.com/cassiusbessa/vision-social-media/domain/core/entities"
	"github.com/google/uuid"
)

type PostRepository interface {
	SavePost(post *entities.ProjectPost) error
	UpdatePost(post *entities.ProjectPost) error
	GetPostByID(postID uuid.UUID) (*entities.ProjectPost, error)
}
