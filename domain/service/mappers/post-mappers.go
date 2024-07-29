package mappers

import (
	"time"

	"github.com/cassiusbessa/vision-social-media/domain/core/entities"
	authorDTO "github.com/cassiusbessa/vision-social-media/domain/service/dtos/author"
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

	uuidPost, err := uuid.Parse(command.ID)
	if err != nil {
		return &entities.ProjectPost{}, errors.NewInvalidArgument("Invalid post ID")
	}

	updatedPost := entities.NewProjectPost(
		entities.PostWithID(uuidPost),
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

func commentEntityToLoadedCommentResponse(comment entities.Comment) commentDTO.LoadCommentResponse {

	var parentID string
	if comment.ParentID == uuid.Nil {
		parentID = ""
	} else {
		parentID = comment.ParentID.String()
	}

	return commentDTO.LoadCommentResponse{
		ID:        comment.ID.String(),
		AuthorID:  comment.UserID.String(),
		PostID:    comment.PostID.String(),
		ParentID:  parentID,
		Content:   comment.Content,
		CreatedAt: comment.CreatedAt.Format(time.RFC3339),
		UpdatedAt: comment.UpdatedAt.Format(time.RFC3339),
	}
}

func reactionEntityToLoadedReactionResponse(reaction entities.Reaction) reactionDTO.LoadReactionResponse {

	var parentID string
	if reaction.ParentID.UUID == uuid.Nil {
		parentID = ""
	} else {
		parentID = reaction.ParentID.UUID.String()
	}

	return reactionDTO.LoadReactionResponse{
		ID:        reaction.ID.String(),
		UserID:    reaction.UserID.String(),
		PostID:    reaction.PostID.String(),
		ParentID:  parentID,
		Type:      reaction.Type,
		CreatedAt: reaction.CreatedAt.Format(time.RFC3339),
	}
}

func AuthorEnitityToLoadedResponse(author entities.Author) authorDTO.AuthorLoadedResponse {
	return authorDTO.AuthorLoadedResponse{
		AuthorID:    author.ID.String(),
		AuthorName:  author.Name,
		AuthorImage: author.Image,
	}
}

func PostEntityToLoadedPostResponse(post entities.ProjectPost) postDTO.LoadedPostResponse {

	comments := make([]commentDTO.LoadCommentResponse, 0)
	for _, comment := range post.Comments {
		comments = append(comments, commentEntityToLoadedCommentResponse(comment))
	}

	reactions := make([]reactionDTO.LoadReactionResponse, 0)
	for _, reaction := range post.Reactions {
		reactions = append(reactions, reactionEntityToLoadedReactionResponse(reaction))
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

func ReactToPostCommandToReactionEntity(command reactionDTO.ReactToPostCommand) (*entities.Reaction, error) {
	uuidPost, err := uuid.Parse(command.PostID)
	if err != nil {
		return &entities.Reaction{}, errors.NewInvalidArgument("Invalid post ID")
	}

	uuidUser, err := uuid.Parse(command.UserID)
	if err != nil {
		return &entities.Reaction{}, errors.NewInvalidArgument("Invalid user ID")
	}

	var reactionType entities.ReactionType

	switch command.Type {
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
	default:
		return &entities.Reaction{}, errors.NewInvalidArgument("Invalid reaction type")
	}

	parentId := uuid.NullUUID{
		UUID:  uuid.Nil,
		Valid: false,
	}

	reaction := entities.NewReaction(
		entities.ReactionWithID(uuid.New()),
		entities.ReactionWithPostID(uuidPost),
		entities.ReactionWithUserID(uuidUser),
		entities.ReactionWithReactionType(reactionType),
		entities.ReactionWithParentID(parentId),
		entities.ReactionWithCreatedAt(time.Now()),
	)

	return reaction, nil
}

func AddCommentToPostCommandToCommentEntity(command commentDTO.AddCommentToPostCommand) (*entities.Comment, error) {
	uuidPost, err := uuid.Parse(command.PostID)
	if err != nil {
		return &entities.Comment{}, errors.NewInvalidArgument("Invalid post ID")
	}

	uuidUser, err := uuid.Parse(command.AuthorID)
	if err != nil {
		return &entities.Comment{}, errors.NewInvalidArgument("Invalid user ID")
	}

	comment := entities.NewComment(
		entities.CommentWithID(uuid.New()),
		entities.CommentWithPostID(uuidPost),
		entities.CommentWithUserID(uuidUser),
		entities.CommentWithContent(command.Content),
		entities.CommentWithCreatedAt(time.Now()),
		entities.CommentWithUpdatedAt(time.Now()),
	)

	return comment, nil
}
