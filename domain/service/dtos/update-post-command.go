package dtos

type UpdatePostCommand struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	RepoLink  string `json:"repoLink"`
	DemoLink  string `json:"demoLink"`
	PostImage string `json:"postImage"`
}

func NewUpdatePostCommand(id, title, content, repoLink, demoLink, postImage string) *UpdatePostCommand {
	return &UpdatePostCommand{
		ID:        id,
		Title:     title,
		Content:   content,
		RepoLink:  repoLink,
		DemoLink:  demoLink,
		PostImage: postImage,
	}
}
