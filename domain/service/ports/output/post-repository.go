package ports

import (
	entities "github.com/cassiusbessa/vision-social-media/domain/core/entities"
)

type PostRepository interface {
	SavePost(post *entities.ProjectPost) error
}
