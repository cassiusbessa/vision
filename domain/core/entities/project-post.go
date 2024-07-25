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

type Option func(*ProjectPost)

func NewProjectPost(opts ...Option) *ProjectPost {

	p := &ProjectPost{}

	for _, opt := range opts {
		opt(p)
	}
	return p
}

func WithID(id uuid.UUID) Option {
	return func(p *ProjectPost) {
		p.ID = id
	}
}

func WithProjectID(projectID uuid.UUID) Option {
	return func(p *ProjectPost) {
		p.ProjectID = projectID
	}
}

func WithAuthorID(authorID uuid.UUID) Option {
	return func(p *ProjectPost) {
		p.AuthorID = authorID
	}
}

func WithTitle(title string) Option {
	return func(p *ProjectPost) {
		p.Title = title
	}
}

func WithContent(content string) Option {
	return func(p *ProjectPost) {
		p.Content = content
	}
}

func WithRepoLink(repoLink string) Option {
	return func(p *ProjectPost) {
		p.RepoLink = repoLink
	}
}

func WithDemoLink(demoLink string) Option {
	return func(p *ProjectPost) {
		p.DemoLink = demoLink
	}
}

func WithPostImage(postImage string) Option {
	return func(p *ProjectPost) {
		p.PostImage = postImage
	}
}

func WithLikeCount(likeCount int) Option {
	return func(p *ProjectPost) {
		p.LikeCount = likeCount
	}
}

func WithCommentCount(commentCount int) Option {
	return func(p *ProjectPost) {
		p.CommentCount = commentCount
	}
}

func WithReactions(reactions []Reaction) Option {
	return func(p *ProjectPost) {
		p.Reactions = reactions
	}
}

func WithComments(comments []Comment) Option {
	return func(p *ProjectPost) {
		p.Comments = comments
	}
}

func WithCreatedAt(createdAt time.Time) Option {
	return func(p *ProjectPost) {
		p.CreatedAt = createdAt
	}
}

func WithUpdatedAt(updatedAt time.Time) Option {
	return func(p *ProjectPost) {
		p.UpdatedAt = updatedAt
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
