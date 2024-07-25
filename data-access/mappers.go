package data

import (
	data "github.com/cassiusbessa/vision-social-media/data-access/sqlc-config"
	"github.com/cassiusbessa/vision-social-media/domain/core/entities"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

func projectEntityToCreateQueryParams(post *entities.ProjectPost) data.CreatePostParams {
	var id, projectId, authorId pgtype.UUID
	var repoLink, demoLink, postImage pgtype.Text
	var createdAt, updatedAt pgtype.Timestamptz

	id.Bytes = [16]byte(post.ID)
	id.Valid = true
	projectId.Bytes = [16]byte(post.ProjectID)
	projectId.Valid = true
	authorId.Bytes = [16]byte(post.AuthorID)
	authorId.Valid = true
	repoLink.String = post.RepoLink
	repoLink.Valid = true
	demoLink.String = post.DemoLink
	demoLink.Valid = true
	postImage.String = post.PostImage
	postImage.Valid = true
	createdAt.Time = post.CreatedAt
	createdAt.Valid = true
	updatedAt.Time = post.UpdatedAt
	updatedAt.Valid = true

	return data.CreatePostParams{
		ID:        id,
		ProjectID: projectId,
		AuthorID:  authorId,
		Title:     post.Title,
		Content:   post.Content,
		RepoLink:  repoLink,
		DemoLink:  demoLink,
		PostImage: postImage,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}
}

func projectEntityToUpdateQueryParams(post *entities.ProjectPost) data.UpdatePostParams {
	var id pgtype.UUID
	var repoLink, demoLink, postImage pgtype.Text
	var updatedAt pgtype.Timestamptz

	id.Bytes = [16]byte(post.ID)
	id.Valid = true
	repoLink.String = post.RepoLink
	repoLink.Valid = true
	demoLink.String = post.DemoLink
	demoLink.Valid = true
	postImage.String = post.PostImage
	postImage.Valid = true
	updatedAt.Time = post.UpdatedAt
	updatedAt.Valid = true

	return data.UpdatePostParams{
		ID:        id,
		Title:     post.Title,
		Content:   post.Content,
		RepoLink:  repoLink,
		DemoLink:  demoLink,
		PostImage: postImage,
		UpdatedAt: updatedAt,
	}
}

func postDBEntityToProjectPost(post data.Post) *entities.ProjectPost {

	var projectPostId, authorId, projectId uuid.UUID

	projectPostId.Scan(post.ID.Bytes)
	authorId.Scan(post.AuthorID.Bytes)
	projectId.Scan(post.ProjectID.Bytes)

	return entities.NewProjectPost(
		entities.WithID(projectPostId),
		entities.WithProjectID(projectId),
		entities.WithAuthorID(authorId),
		entities.WithTitle(post.Title),
		entities.WithContent(post.Content),
		entities.WithRepoLink(post.RepoLink.String),
		entities.WithDemoLink(post.DemoLink.String),
		entities.WithPostImage(post.PostImage.String),
		entities.WithLikeCount(int(post.LikeCount.Int32)),
		entities.WithCommentCount(int(post.CommentCount.Int32)),
		entities.WithReactions([]entities.Reaction{}),
		entities.WithComments([]entities.Comment{}),
		entities.WithCreatedAt(post.CreatedAt.Time),
		entities.WithUpdatedAt(post.UpdatedAt.Time),
	)
}
