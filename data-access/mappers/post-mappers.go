package mappers

import (
	"database/sql"

	data "github.com/cassiusbessa/vision-social-media/data-access/sqlc-config"
	"github.com/cassiusbessa/vision-social-media/domain/core/entities"
)

func PostEntityToCreateQueryParams(post *entities.ProjectPost) data.CreatePostParams {
	var repoLink, demoLink, postImage sql.NullString

	if post.RepoLink != "" {
		repoLink.String = post.RepoLink
		repoLink.Valid = true
	}
	if post.DemoLink != "" {
		demoLink.String = post.DemoLink
		demoLink.Valid = true
	}
	if post.PostImage != "" {
		postImage.String = post.PostImage
		postImage.Valid = true
	}

	return data.CreatePostParams{
		ID:        post.ID,
		ProjectID: post.ProjectID,
		AuthorID:  post.Author.ID,
		Title:     post.Title,
		Content:   post.Content,
		RepoLink:  repoLink,
		DemoLink:  demoLink,
		PostImage: postImage,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
}

func PostEntityToUpdateQueryParams(post *entities.ProjectPost) data.UpdatePostParams {
	var repoLink, demoLink, postImage sql.NullString

	if post.RepoLink != "" {
		repoLink.String = post.RepoLink
		repoLink.Valid = true
	}
	if post.DemoLink != "" {
		demoLink.String = post.DemoLink
		demoLink.Valid = true
	}
	if post.PostImage != "" {
		postImage.String = post.PostImage
		postImage.Valid = true
	}

	return data.UpdatePostParams{
		ID:        post.ID,
		Title:     post.Title,
		Content:   post.Content,
		RepoLink:  repoLink,
		DemoLink:  demoLink,
		PostImage: postImage,
		UpdatedAt: post.UpdatedAt,
	}
}

func PostDBEntityToProjectPost(post data.Post, comments []data.Comment, reactions []data.Reaction) *entities.ProjectPost {

	entityComments := make([]entities.Comment, 0, len(comments))
	for _, comment := range comments {
		entityComments = append(entityComments, *CommentDbEntityToComment(comment))
	}

	entityReactions := make([]entities.Reaction, 0, len(reactions))
	for _, reaction := range reactions {
		entityReactions = append(entityReactions, *ReactionDbEntityToReaction(reaction))
	}

	return entities.NewProjectPost(
		entities.PostWithID(post.ID),
		entities.PostWithProjectID(post.ProjectID),
		entities.PostWithAuthorID(post.AuthorID),
		entities.PostWithTitle(post.Title),
		entities.PostWithContent(post.Content),
		entities.PostWithRepoLink(post.RepoLink.String),
		entities.PostWithDemoLink(post.DemoLink.String),
		entities.PostWithPostImage(post.PostImage.String),
		entities.PostWithLikeCount(int(post.LikeCount)),
		entities.PostWithCommentCount(int(post.CommentCount)),
		entities.PostWithReactions(entityReactions),
		entities.PostWithComments(entityComments),
		entities.PostWithCreatedAt(post.CreatedAt),
		entities.WithUpdatedAt(post.UpdatedAt),
	)
}

func LoadOrderedPostRowToProjectPosts(post data.LoadOrderedPostsRow) *entities.ProjectPost {

	return entities.NewProjectPost(
		entities.PostWithID(post.PostID),
		entities.PostWithProjectID(post.ProjectID),
		entities.PostWithAuthorID(post.AuthorID),
		entities.PostWithAuthorImage(post.AuthorImage.String),
		entities.PostWithAuthorName(post.AuthorName),
		entities.PostWithTitle(post.Title),
		entities.PostWithContent(post.PostContent),
		entities.PostWithRepoLink(post.RepoLink.String),
		entities.PostWithDemoLink(post.DemoLink.String),
		entities.PostWithPostImage(post.PostImage.String),
		entities.PostWithLikeCount(int(post.LikeCount)),
		entities.PostWithCommentCount(int(post.CommentCount)),
		entities.PostWithReactions([]entities.Reaction{}),
		entities.PostWithComments([]entities.Comment{}),
		entities.PostWithCreatedAt(post.PostCreatedAt),
		entities.WithUpdatedAt(post.PostUpdatedAt),
	)
}
