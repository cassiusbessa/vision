package dtos

type AddCommentToPostCommand struct {
	PostID   string `json:"post_id"`
	ParentID string `json:"parent_id"`
	AuthorID string `json:"author_id"`
	Content  string `json:"content"`
}
