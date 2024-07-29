package dtos

type LoadCommentResponse struct {
	ID        string `json:"id"`
	PostID    string `json:"post_id"`
	ParentID  string `json:"parent_id"`
	AuthorID  string `json:"author_id"`
	Content   string `json:"content"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
