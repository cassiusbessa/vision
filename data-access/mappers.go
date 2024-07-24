package data

import (
	data "github.com/cassiusbessa/vision-social-media/data-access/sqlc-config"
	"github.com/cassiusbessa/vision-social-media/domain/core/entities"
	"github.com/jackc/pgx/v5/pgtype"
)

func projectEntityToQueryParams(post *entities.ProjectPost) data.CreatePostParams {
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
