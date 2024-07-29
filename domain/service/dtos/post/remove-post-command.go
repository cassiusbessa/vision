package dtos

type RemovePostCommand struct {
	PostID string `json:"post_id"`
	UserID string `json:"user_id"`
}
