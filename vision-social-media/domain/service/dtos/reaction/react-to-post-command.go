package dtos

type ReactToPostCommand struct {
	PostID string `json:"post_id"`
	UserID string `json:"user_id"`
	Type   string `json:"type"`
}
