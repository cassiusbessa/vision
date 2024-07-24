package ports

import (
	dto "github.com/cassiusbessa/vision-social-media/domain/service/dtos"
)

type PostService interface {
	CreatePost(command *dto.CreatePostCommand) error
}
