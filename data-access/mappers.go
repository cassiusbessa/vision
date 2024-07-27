package data

import (
	"database/sql"

	data "github.com/cassiusbessa/vision-social-media/data-access/sqlc-config"
	"github.com/cassiusbessa/vision-social-media/domain/core/entities"
)

func projectEntityToCreateQueryParams(post *entities.ProjectPost) data.CreatePostParams {
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
		AuthorID:  post.AuthorID,
		Title:     post.Title,
		Content:   post.Content,
		RepoLink:  repoLink,
		DemoLink:  demoLink,
		PostImage: postImage,
		CreatedAt: post.CreatedAt,
		UpdatedAt: post.UpdatedAt,
	}
}

func projectEntityToUpdateQueryParams(post *entities.ProjectPost) data.UpdatePostParams {
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

func postDBEntityToProjectPost(post data.Post, comments []data.Comment, reactions []data.Reaction) *entities.ProjectPost {

	entityComments := make([]entities.Comment, 0, len(comments))
	for _, comment := range comments {
		entityComments = append(entityComments, *commentDbEntityToComment(comment))
	}

	entityReactions := make([]entities.Reaction, 0, len(reactions))
	for _, reaction := range reactions {
		entityReactions = append(entityReactions, *reactionDbEntityToReaction(reaction))
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

func loadOrderedPostRowToProjectComment(post data.LoadOrderedPostsRow) *entities.Comment {

	return entities.NewComment(
		entities.CommentWithID(post.CommentID.UUID),
		entities.CommentWithPostID(post.PostID),
		entities.CommentWithParentID(post.CommentParentID.UUID),
		entities.CommentWithUserID(post.CommentUserID.UUID),
		entities.CommentWithContent(post.CommentContent.String),
		entities.CommentWithReactions([]entities.Reaction{}),
		entities.CommentWithCreatedAt(post.CommentCreatedAt.Time),
		entities.CommentWithUpdatedAt(post.CommentUpdatedAt.Time),
	)
}

func loadOrderedPostRowToProjectReaction(post data.LoadOrderedPostsRow) *entities.Reaction {

	var reactionType entities.ReactionType

	switch post.ReactionType.String {
	case "Like":
		reactionType = entities.Like
	case "Dislike":
		reactionType = entities.Dislike
	case "Love":
		reactionType = entities.Love
	case "Wow":
		reactionType = entities.Wow
	case "Angry":
		reactionType = entities.Angry
	}

	return entities.NewReaction(
		entities.ReactionWithID(post.ReactionID.UUID),
		entities.ReactionWithPostID(post.PostID),
		entities.ReactionWithUserID(post.ReactionUserID.UUID),
		entities.ReactionWithReactionType(reactionType),
	)
}

func loadOrderedPostRowToProjectPosts(post data.LoadOrderedPostsRow) *entities.ProjectPost {

	return entities.NewProjectPost(
		entities.PostWithID(post.PostID),
		entities.PostWithProjectID(post.ProjectID),
		entities.PostWithAuthorID(post.AuthorID),
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

func reactionEntityToCreateQueryParams(reaction *entities.Reaction) data.CreateReactionParams {
	return data.CreateReactionParams{
		ID:           reaction.ID,
		PostID:       reaction.PostID,
		CommentID:    reaction.ParentID,
		UserID:       reaction.UserID,
		ReactionType: string(reaction.Type),
	}
}

func reactionDbEntityToReaction(reaction data.Reaction) *entities.Reaction {

	var reactionType entities.ReactionType

	switch reaction.ReactionType {
	case "Like":
		reactionType = entities.Like
	case "Dislike":
		reactionType = entities.Dislike
	case "Love":
		reactionType = entities.Love
	case "Wow":
		reactionType = entities.Wow
	case "Angry":
		reactionType = entities.Angry
	}

	return entities.NewReaction(
		entities.ReactionWithID(reaction.ID),
		entities.ReactionWithPostID(reaction.PostID),
		entities.ReactionWithUserID(reaction.UserID),
		entities.ReactionWithReactionType(reactionType),
	)
}

func commentDbEntityToComment(comment data.Comment) *entities.Comment {

	return entities.NewComment(
		entities.CommentWithID(comment.ID),
		entities.CommentWithPostID(comment.PostID),
		entities.CommentWithParentID(comment.ParentID.UUID),
		entities.CommentWithUserID(comment.UserID),
		entities.CommentWithContent(comment.Content),
		entities.CommentWithReactions([]entities.Reaction{}),
		entities.CommentWithCreatedAt(comment.CreatedAt.Time),
		entities.CommentWithUpdatedAt(comment.UpdatedAt.Time),
	)
}
