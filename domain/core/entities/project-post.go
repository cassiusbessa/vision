package entities

import (
	"time"

	"github.com/google/uuid"
)

type ProjectPost struct {
	ID             uuid.UUID
	ProjectID      uuid.UUID
	AuthorID       uuid.UUID
	Title          string
	Content        string
	RepoLink       string
	DemoLink       string
	PostImage      string
	LikeCount      int
	CommentCount   int
	Reactions      []Reaction
	Comments       []Comment
	CreatedAt      time.Time
	UpdatedAt      time.Time
	FailureMessage []string
}

func NewProjectPost(projectID, authorID uuid.UUID, title, content, repoLink, demoLink, postImage string) *ProjectPost {
	return &ProjectPost{
		ID:           uuid.New(),
		ProjectID:    projectID,
		AuthorID:     authorID,
		Title:        title,
		Content:      content,
		RepoLink:     repoLink,
		DemoLink:     demoLink,
		PostImage:    postImage,
		LikeCount:    0,
		CommentCount: 0,
		Reactions:    []Reaction{},
		Comments:     []Comment{},
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}

func (p *ProjectPost) Validate() {
	if p.Title == "" || len(p.Title) > 255 {
		p.FailureMessage = append(p.FailureMessage, "Project post title cannot be empty or more than 255 characters")
	}
	if len(p.Content) > 2000 {
		p.FailureMessage = append(p.FailureMessage, "Project post content cannot be empty or more than 1000 characters")
	}
	if p.RepoLink != "" && len(p.RepoLink) > 255 {
		p.FailureMessage = append(p.FailureMessage, "Project post repository link cannot be more than 255 characters")
	}
	if p.DemoLink != "" && len(p.DemoLink) > 255 {
		p.FailureMessage = append(p.FailureMessage, "Project post demo link cannot be more than 255 characters")
	}
	if p.PostImage != "" && len(p.PostImage) > 255 {
		p.FailureMessage = append(p.FailureMessage, "Project post image link cannot be more than 255 characters")
	}
}
