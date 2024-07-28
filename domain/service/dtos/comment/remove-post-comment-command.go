package dtos

type RemovePostCommentCommand struct {
	PostID    string `json:"post_id"`
	CommentID string `json:"comment_id"`
	UserID    string `json:"user_id"`
}
