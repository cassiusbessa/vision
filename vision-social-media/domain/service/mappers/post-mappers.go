package mappers

import (
	"time"

	"github.com/cassiusbessa/vision-social-media/domain/core/entities"
	commentDTO "github.com/cassiusbessa/vision-social-media/domain/service/dtos/comment"
	postDTO "github.com/cassiusbessa/vision-social-media/domain/service/dtos/post"
	reactionDTO "github.com/cassiusbessa/vision-social-media/domain/service/dtos/reaction"
	"github.com/cassiusbessa/vision-social-media/domain/service/errors"
	"github.com/google/uuid"
)

func CreatePostCommandToPostEntity(command postDTO.CreatePostCommand) (*entities.ProjectPost, error) {
	uuidProject, err := uuid.Parse(command.ProjectID)
	if err != nil {
		return &entities.ProjectPost{}, errors.NewInvalidArgument("Invalid project ID")
	}

	uuidUser, err := uuid.Parse(command.AuthorID)
	if err != nil {
		return &entities.ProjectPost{}, errors.NewInvalidArgument("Invalid user ID")
	}
	post := entities.NewProjectPost(
		entities.PostWithID(uuid.New()),
		entities.PostWithProjectID(uuidProject),
		entities.PostWithAuthorID(uuidUser),
		entities.PostWithTitle(command.Title),
		entities.PostWithContent(command.Content),
		entities.PostWithRepoLink(command.RepoLink),
		entities.PostWithDemoLink(command.DemoLink),
		entities.PostWithPostImage(command.PostImage),
		entities.PostWithLikeCount(0),
		entities.PostWithCommentCount(0),
		entities.PostWithReactions([]entities.Reaction{}),
		entities.PostWithComments([]entities.Comment{}),
		entities.PostWithCreatedAt(time.Now()),
		entities.WithUpdatedAt(time.Now()),
	)
	return post, nil
}

func UpdatePostCommandToPostEntity(command postDTO.UpdatePostCommand, post entities.ProjectPost) (*entities.ProjectPost, error) {

	updatedPost := entities.NewProjectPost(
		entities.PostWithID(post.ID),
		entities.PostWithProjectID(post.ProjectID),
		entities.PostWithAuthorID(post.Author.ID),
		entities.PostWithTitle(command.Title),
		entities.PostWithContent(command.Content),
		entities.PostWithRepoLink(command.RepoLink),
		entities.PostWithDemoLink(command.DemoLink),
		entities.PostWithPostImage(command.PostImage),
		entities.PostWithLikeCount(post.LikeCount),
		entities.PostWithCommentCount(post.CommentCount),
		entities.PostWithReactions(post.Reactions),
		entities.PostWithComments(post.Comments),
		entities.PostWithCreatedAt(post.CreatedAt),
		entities.WithUpdatedAt(time.Now()),
	)

	return updatedPost, nil
}

func PostEntityToLoadedPostResponse(post entities.ProjectPost) postDTO.LoadedPostResponse {

	comments := make([]commentDTO.LoadedCommentResponse, 0)
	for _, comment := range post.Comments {
		comments = append(comments, CommentEntityToLoadedCommentResponse(comment))
	}

	reactions := make([]reactionDTO.LoadReactionResponse, 0)
	for _, reaction := range post.Reactions {
		reactions = append(reactions, ReactionEntityToLoadedReactionResponse(reaction))
	}

	return postDTO.LoadedPostResponse{
		ID:           post.ID.String(),
		ProjectID:    post.ProjectID.String(),
		Author:       AuthorEnitityToLoadedResponse(post.Author),
		Title:        post.Title,
		Content:      post.Content,
		RepoLink:     post.RepoLink,
		DemoLink:     post.DemoLink,
		PostImage:    post.PostImage,
		LikeCount:    post.LikeCount,
		CommentCount: post.CommentCount,
		Comments:     comments,
		Reactions:    reactions,
		CreatedAt:    post.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    post.UpdatedAt.Format(time.RFC3339),
	}
}
