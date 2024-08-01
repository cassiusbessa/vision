package dtos

type UpdatePostCommand struct {
	ProjectID string `json:"project_id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	RepoLink  string `json:"repo_link"`
	DemoLink  string `json:"demo_link"`
	PostImage string `json:"post_image"`
}

func NewUpdatePostCommand(projectID, title, content, repoLink, demoLink, postImage string) *UpdatePostCommand {
	return &UpdatePostCommand{
		ProjectID: projectID,
		Title:     title,
		Content:   content,
		RepoLink:  repoLink,
		DemoLink:  demoLink,
		PostImage: postImage,
	}
}
