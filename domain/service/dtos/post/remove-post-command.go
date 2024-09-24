package dtos

type RemovePostCommand struct {
	ProjectID string `json:"project_id"`
	UserID    string `json:"user_id"`
}
