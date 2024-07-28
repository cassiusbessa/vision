package dtos

import (
	dtosComment "github.com/cassiusbessa/vision-social-media/domain/service/dtos/comment"
	dtosReaction "github.com/cassiusbessa/vision-social-media/domain/service/dtos/reaction"
)

type LoadedPostResponse struct {
	ID           string                              `json:"id"`
	ProjectID    string                              `json:"project_id"`
	AuthorID     string                              `json:"author_id"`
	Title        string                              `json:"title"`
	Content      string                              `json:"content"`
	RepoLink     string                              `json:"repo_link"`
	DemoLink     string                              `json:"demo_link"`
	PostImage    string                              `json:"post_image"`
	LikeCount    int                                 `json:"like_count"`
	CommentCount int                                 `json:"comment_count"`
	Comments     []dtosComment.LoadCommentResponse   `json:"comments"`
	Reactions    []dtosReaction.LoadReactionResponse `json:"reactions"`
	CreatedAt    string                              `json:"created_at"`
	UpdatedAt    string                              `json:"updated_at"`
}
