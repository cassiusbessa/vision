package mappers

import (
	"time"

	"github.com/cassiusbessa/vision-social-media/domain/core/entities"
	"github.com/cassiusbessa/vision-social-media/domain/service/dtos"
	"github.com/cassiusbessa/vision-social-media/domain/service/errors"
	"github.com/google/uuid"
)

func CreatePostCommandToPostEntity(command dtos.CreatePostCommand) (*entities.ProjectPost, error) {
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

func UpdatePostCommandToPostEntity(command dtos.UpdatePostCommand, post entities.ProjectPost) (*entities.ProjectPost, error) {

	uuidPost, err := uuid.Parse(command.ID)
	if err != nil {
		return &entities.ProjectPost{}, errors.NewInvalidArgument("Invalid post ID")
	}

	updatedPost := entities.NewProjectPost(
		entities.PostWithID(uuidPost),
		entities.PostWithProjectID(post.ProjectID),
		entities.PostWithAuthorID(post.AuthorID),
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

func commentEntityToLoadedCommentResponse(comment entities.Comment) dtos.LoadCommentResponse {
	return dtos.LoadCommentResponse{
		ID:        comment.ID.String(),
		AuthorID:  comment.UserID.String(),
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt.Format(time.RFC3339),
		UpdatedAt: comment.UpdatedAt.Format(time.RFC3339),
	}
}

func PostEntityToLoadedPostResponse(post entities.ProjectPost) dtos.LoadedPostResponse {

	comments := make([]dtos.LoadCommentResponse, 0)
	for _, comment := range post.Comments {
		comments = append(comments, commentEntityToLoadedCommentResponse(comment))
	}
	return dtos.LoadedPostResponse{
		ID:           post.ID.String(),
		ProjectID:    post.ProjectID.String(),
		AuthorID:     post.AuthorID.String(),
		Title:        post.Title,
		Content:      post.Content,
		RepoLink:     post.RepoLink,
		DemoLink:     post.DemoLink,
		PostImage:    post.PostImage,
		LikeCount:    post.LikeCount,
		CommentCount: post.CommentCount,
		Comments:     comments,
		CreatedAt:    post.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    post.UpdatedAt.Format(time.RFC3339),
	}
}
