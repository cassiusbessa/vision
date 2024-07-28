package ports

import (
	"github.com/cassiusbessa/vision-social-media/domain/service/dtos"
)

type PostService interface {
	CreatePost(command *dtos.CreatePostCommand) (dtos.CreatedPostResponse, error)
	UpdatePost(command *dtos.UpdatePostCommand) (dtos.UpdatedPostResponse, error)
	LoadOrderedPosts(query *dtos.LoadOrderedPostsQuery) ([]dtos.LoadedPostResponse, error)
	ReactToPost(command *dtos.ReactToPostCommand) (dtos.ReactToPostResponse, error)
	RemovePostReaction(command *dtos.RemovePostReactionCommand) (dtos.RemovePostReactionResponse, error)
}
