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

type PostOption func(*ProjectPost)

func NewProjectPost(opts ...PostOption) *ProjectPost {

	p := &ProjectPost{}

	for _, opt := range opts {
		opt(p)
	}
	return p
}

func PostWithID(id uuid.UUID) PostOption {
	return func(p *ProjectPost) {
		p.ID = id
	}
}

func PostWithProjectID(projectID uuid.UUID) PostOption {
	return func(p *ProjectPost) {
		p.ProjectID = projectID
	}
}

func PostWithAuthorID(authorID uuid.UUID) PostOption {
	return func(p *ProjectPost) {
		p.AuthorID = authorID
	}
}

func PostWithTitle(title string) PostOption {
	return func(p *ProjectPost) {
		p.Title = title
	}
}

func PostWithContent(content string) PostOption {
	return func(p *ProjectPost) {
		p.Content = content
	}
}

func PostWithRepoLink(repoLink string) PostOption {
	return func(p *ProjectPost) {
		p.RepoLink = repoLink
	}
}

func PostWithDemoLink(demoLink string) PostOption {
	return func(p *ProjectPost) {
		p.DemoLink = demoLink
	}
}

func PostWithPostImage(postImage string) PostOption {
	return func(p *ProjectPost) {
		p.PostImage = postImage
	}
}

func PostWithLikeCount(likeCount int) PostOption {
	return func(p *ProjectPost) {
		p.LikeCount = likeCount
	}
}

func PostWithCommentCount(commentCount int) PostOption {
	return func(p *ProjectPost) {
		p.CommentCount = commentCount
	}
}

func PostWithReactions(reactions []Reaction) PostOption {
	return func(p *ProjectPost) {
		p.Reactions = reactions
	}
}

func PostWithComments(comments []Comment) PostOption {
	return func(p *ProjectPost) {
		p.Comments = comments
	}
}

func PostWithCreatedAt(createdAt time.Time) PostOption {
	return func(p *ProjectPost) {
		p.CreatedAt = createdAt
	}
}

func WithUpdatedAt(updatedAt time.Time) PostOption {
	return func(p *ProjectPost) {
		p.UpdatedAt = updatedAt
	}
}

func (p *ProjectPost) AddReaction(reaction *Reaction) {
	p.Reactions = append(p.Reactions, *reaction)
	p.LikeCount++
}

func (p *ProjectPost) AddComment(comment *Comment) {
	p.Comments = append(p.Comments, *comment)
	p.CommentCount++
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
