package dtos

type AddCommentToPostResponse struct {
	CommentID string `json:"comment_id"`
	Message   string `json:"message"`
}
