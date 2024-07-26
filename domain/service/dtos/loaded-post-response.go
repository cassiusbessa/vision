package dtos

type LoadedPostResponse struct {
	ID           string                `json:"id"`
	ProjectID    string                `json:"project_id"`
	AuthorID     string                `json:"author_id"`
	Title        string                `json:"title"`
	Content      string                `json:"content"`
	RepoLink     string                `json:"repo_link"`
	DemoLink     string                `json:"demo_link"`
	PostImage    string                `json:"post_image"`
	LikeCount    int                   `json:"like_count"`
	CommentCount int                   `json:"comment_count"`
	Comments     []LoadCommentResponse `json:"comments"`
	CreatedAt    string                `json:"created_at"`
	UpdatedAt    string                `json:"updated_at"`
}
