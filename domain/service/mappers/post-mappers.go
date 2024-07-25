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
		entities.WithID(uuid.New()),
		entities.WithProjectID(uuidProject),
		entities.WithAuthorID(uuidUser),
		entities.WithTitle(command.Title),
		entities.WithContent(command.Content),
		entities.WithRepoLink(command.RepoLink),
		entities.WithDemoLink(command.DemoLink),
		entities.WithPostImage(command.PostImage),
		entities.WithLikeCount(0),
		entities.WithCommentCount(0),
		entities.WithReactions([]entities.Reaction{}),
		entities.WithComments([]entities.Comment{}),
		entities.WithCreatedAt(time.Now()),
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
		entities.WithID(uuidPost),
		entities.WithProjectID(post.ProjectID),
		entities.WithAuthorID(post.AuthorID),
		entities.WithTitle(command.Title),
		entities.WithContent(command.Content),
		entities.WithRepoLink(command.RepoLink),
		entities.WithDemoLink(command.DemoLink),
		entities.WithPostImage(command.PostImage),
		entities.WithLikeCount(post.LikeCount),
		entities.WithCommentCount(post.CommentCount),
		entities.WithReactions(post.Reactions),
		entities.WithComments(post.Comments),
		entities.WithCreatedAt(post.CreatedAt),
		entities.WithUpdatedAt(time.Now()),
	)

	return updatedPost, nil
}
