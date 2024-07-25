package ports

import (
	"github.com/cassiusbessa/vision-social-media/domain/service/dtos"
	dto "github.com/cassiusbessa/vision-social-media/domain/service/dtos"
)

type PostService interface {
	CreatePost(command *dto.CreatePostCommand) (dtos.CreatedPostResponse, error)
}
