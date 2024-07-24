package dtos

type CreatePostCommand struct {
	ProjectID string `json:"project_id"`
	AuthorID  string `json:"author_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	RepoLink  string `json:"repo_link"`
	DemoLink  string `json:"demo_link"`
	PostImage string `json:"post_image"`
}

func NewCreatePostCommand(projectID, authorID, title, content, repoLink, demoLink, postImage string) *CreatePostCommand {
	return &CreatePostCommand{
		ProjectID: projectID,
		AuthorID:  authorID,
		Title:     title,
		Content:   content,
		RepoLink:  repoLink,
		DemoLink:  demoLink,
		PostImage: postImage,
	}
}
