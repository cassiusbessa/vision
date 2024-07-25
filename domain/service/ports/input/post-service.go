package ports

import (
	"github.com/cassiusbessa/vision-social-media/domain/service/dtos"
)

type PostService interface {
	CreatePost(command *dtos.CreatePostCommand) (dtos.CreatedPostResponse, error)
	UpdatePost(command *dtos.UpdatePostCommand) (dtos.UpdatedPostResponse, error)
}
