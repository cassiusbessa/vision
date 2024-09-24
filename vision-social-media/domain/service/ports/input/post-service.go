package ports

import (
	commentDTO "github.com/cassiusbessa/vision-social-media/domain/service/dtos/comment"
	postDTO "github.com/cassiusbessa/vision-social-media/domain/service/dtos/post"
	reactionDTO "github.com/cassiusbessa/vision-social-media/domain/service/dtos/reaction"
)

type PostService interface {
	CreatePost(command *postDTO.CreatePostCommand) (postDTO.CreatedPostResponse, error)
	UpdatePost(command *postDTO.UpdatePostCommand) (postDTO.UpdatedPostResponse, error)
	DeletePost(command *postDTO.RemovePostCommand) (postDTO.RemovedPostResponse, error)
	LoadOrderedPosts(query *postDTO.LoadOrderedPostsQuery) ([]postDTO.LoadedPostResponse, error)
	ReactToPost(command *reactionDTO.ReactToPostCommand) (reactionDTO.ReactToPostResponse, error)
	RemovePostReaction(command *reactionDTO.RemovePostReactionCommand) (reactionDTO.RemovePostReactionResponse, error)
	LoadPostReactionsByPostID(query *reactionDTO.LoadOrderedReactionsQuery) ([]reactionDTO.LoadReactionResponse, error)
	AddCommentToPost(command *commentDTO.AddCommentToPostCommand) (commentDTO.AddCommentToPostResponse, error)
	RemovePostComment(command *commentDTO.RemovePostCommentCommand) (commentDTO.RemovePostCommentResponse, error)
	LoadPostCommentsByPostID(query *commentDTO.LoadOrderedCommentsQuery) ([]commentDTO.LoadedCommentResponse, error)
}
