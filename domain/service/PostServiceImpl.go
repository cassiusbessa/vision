package service

import (
	"strings"

	commentDTO "github.com/cassiusbessa/vision-social-media/domain/service/dtos/comment"
	postDTO "github.com/cassiusbessa/vision-social-media/domain/service/dtos/post"
	reactionDTO "github.com/cassiusbessa/vision-social-media/domain/service/dtos/reaction"
	errors "github.com/cassiusbessa/vision-social-media/domain/service/errors"
	"github.com/cassiusbessa/vision-social-media/domain/service/mappers"
	outputPorts "github.com/cassiusbessa/vision-social-media/domain/service/ports/output"
	"github.com/google/uuid"
)

type PostService struct {
	postRepo outputPorts.PostRepository
}

func NewPostService(postRepo outputPorts.PostRepository) *PostService {
	return &PostService{
		postRepo: postRepo,
	}
}

func (service *PostService) CreatePost(command *postDTO.CreatePostCommand) (postDTO.CreatedPostResponse, error) {

	post, err := mappers.CreatePostCommandToPostEntity(*command)
	if err != nil {
		return postDTO.CreatedPostResponse{}, err
	}

	post.Validate()
	if len(post.FailureMessage) > 0 {
		return postDTO.CreatedPostResponse{}, errors.NewValidationError(strings.Join(post.FailureMessage, ", "))
	}

	err = service.postRepo.SavePost(post)
	if err != nil {
		return postDTO.CreatedPostResponse{}, err
	}

	return postDTO.CreatedPostResponse{
		ID:      post.ID.String(),
		Message: "Post created",
	}, nil
}

func (service *PostService) DeletePost(command *postDTO.RemovePostCommand) (postDTO.RemovedPostResponse, error) {

	uuidPost, err := uuid.Parse(command.PostID)
	if err != nil {
		return postDTO.RemovedPostResponse{}, errors.NewInvalidArgument("Invalid post ID")
	}

	removed, err := service.postRepo.RemovePost(uuidPost)
	if err != nil {
		return postDTO.RemovedPostResponse{}, err
	}
	if !removed {
		return postDTO.RemovedPostResponse{}, errors.NewResourceNotFound("Post not found")
	}

	return postDTO.RemovedPostResponse{
		Message: "Post removed",
	}, nil
}

func (service *PostService) UpdatePost(command *postDTO.UpdatePostCommand) (postDTO.UpdatedPostResponse, error) {

	uuidPost, err := uuid.Parse(command.ID)
	if err != nil {
		return postDTO.UpdatedPostResponse{}, errors.NewInvalidArgument("Invalid post ID")
	}

	post, err := service.postRepo.GetPostByID(uuidPost)
	if err != nil {
		return postDTO.UpdatedPostResponse{}, err
	}
	if post == nil {
		return postDTO.UpdatedPostResponse{}, errors.NewResourceNotFound("Post not found")
	}

	updatedPost, err := mappers.UpdatePostCommandToPostEntity(*command, *post)
	if err != nil {
		return postDTO.UpdatedPostResponse{}, err
	}

	post.Validate()
	if len(post.FailureMessage) > 0 {
		return postDTO.UpdatedPostResponse{}, errors.NewValidationError(strings.Join(post.FailureMessage, ", "))
	}

	err = service.postRepo.UpdatePost(updatedPost)
	if err != nil {
		return postDTO.UpdatedPostResponse{}, err
	}

	return postDTO.UpdatedPostResponse{
		ID:      post.ID.String(),
		Message: "Post updated",
	}, nil
}

func (service *PostService) LoadOrderedPosts(query *postDTO.LoadOrderedPostsQuery) ([]postDTO.LoadedPostResponse, error) {

	posts, err := service.postRepo.LoadOrderedPosts()
	if err != nil {
		return nil, err
	}

	loadedPosts := make([]postDTO.LoadedPostResponse, 0)
	for _, post := range posts {
		loadedPosts = append(loadedPosts, mappers.PostEntityToLoadedPostResponse(post))
	}

	return loadedPosts, nil
}

func (service *PostService) ReactToPost(command *reactionDTO.ReactToPostCommand) (reactionDTO.ReactToPostResponse, error) {

	uuidPost, err := uuid.Parse(command.PostID)
	if err != nil {
		return reactionDTO.ReactToPostResponse{}, errors.NewInvalidArgument("Invalid post ID")
	}

	post, err := service.postRepo.GetPostByID(uuidPost)
	if err != nil {
		return reactionDTO.ReactToPostResponse{}, err
	}
	if post == nil {
		return reactionDTO.ReactToPostResponse{}, errors.NewResourceNotFound("Post not found")
	}
	for _, reaction := range post.Reactions {
		if reaction.UserID == uuid.MustParse(command.UserID) {
			return reactionDTO.ReactToPostResponse{}, errors.NewResourceAlreadyExists("User already reacted to this post")
		}
	}

	reaction, err := mappers.ReactToPostCommandToReactionEntity(*command)
	if err != nil {
		return reactionDTO.ReactToPostResponse{}, err
	}

	reaction.Validate()
	if len(reaction.FailureMessage) > 0 {
		return reactionDTO.ReactToPostResponse{}, errors.NewValidationError(strings.Join(reaction.FailureMessage, ", "))
	}

	err = service.postRepo.AddReactionToPost(reaction)
	if err != nil {
		return reactionDTO.ReactToPostResponse{}, err
	}

	return reactionDTO.ReactToPostResponse{
		ID:      reaction.ID.String(),
		Message: "Reaction saved",
	}, nil
}

func (service *PostService) RemovePostReaction(command *reactionDTO.RemovePostReactionCommand) (reactionDTO.RemovePostReactionResponse, error) {
	uuidReaction, err := uuid.Parse(command.ReactionID)
	if err != nil {
		return reactionDTO.RemovePostReactionResponse{}, errors.NewInvalidArgument("Invalid reaction ID")
	}

	uuidPost, err := uuid.Parse(command.PostID)
	if err != nil {
		return reactionDTO.RemovePostReactionResponse{}, errors.NewInvalidArgument("Invalid post ID")
	}

	removed, err := service.postRepo.RemovePostReaction(uuidReaction, uuidPost)
	if err != nil {
		return reactionDTO.RemovePostReactionResponse{}, err
	}
	if !removed {
		return reactionDTO.RemovePostReactionResponse{}, errors.NewResourceNotFound("Reaction not found")
	}

	return reactionDTO.RemovePostReactionResponse{
		Message: "Reaction removed",
	}, nil
}

func (service *PostService) AddCommentToPost(command *commentDTO.AddCommentToPostCommand) (commentDTO.AddCommentToPostResponse, error) {

	uuidPost, err := uuid.Parse(command.PostID)
	if err != nil {
		return commentDTO.AddCommentToPostResponse{}, errors.NewInvalidArgument("Invalid post ID")
	}

	post, err := service.postRepo.GetPostByID(uuidPost)
	if err != nil {
		return commentDTO.AddCommentToPostResponse{}, err
	}
	if post == nil {
		return commentDTO.AddCommentToPostResponse{}, errors.NewResourceNotFound("Post not found")
	}

	comment, err := mappers.AddCommentToPostCommandToCommentEntity(*command)
	if err != nil {
		return commentDTO.AddCommentToPostResponse{}, err
	}

	comment.Validate()
	if len(comment.FailureMessage) > 0 {
		return commentDTO.AddCommentToPostResponse{}, errors.NewValidationError(strings.Join(comment.FailureMessage, ", "))
	}

	err = service.postRepo.AddCommentToPost(comment)
	if err != nil {
		return commentDTO.AddCommentToPostResponse{}, err
	}

	return commentDTO.AddCommentToPostResponse{
		CommentID: comment.ID.String(),
		Message:   "Comment saved",
	}, nil
}

func (service *PostService) RemovePostComment(command *commentDTO.RemovePostCommentCommand) (commentDTO.RemovePostCommentResponse, error) {
	uuidComment, err := uuid.Parse(command.CommentID)
	if err != nil {
		return commentDTO.RemovePostCommentResponse{}, errors.NewInvalidArgument("Invalid comment ID")
	}

	uuidPost, err := uuid.Parse(command.PostID)
	if err != nil {
		return commentDTO.RemovePostCommentResponse{}, errors.NewInvalidArgument("Invalid post ID")
	}

	removed, err := service.postRepo.RemovePostComment(uuidComment, uuidPost)
	if err != nil {
		return commentDTO.RemovePostCommentResponse{}, err
	}
	if !removed {
		return commentDTO.RemovePostCommentResponse{}, errors.NewResourceNotFound("Comment not found")
	}

	return commentDTO.RemovePostCommentResponse{
		Message: "Comment removed",
	}, nil
}
