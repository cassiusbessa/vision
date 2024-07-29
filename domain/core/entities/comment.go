package entities

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	ID             uuid.UUID
	PostID         uuid.UUID
	ParentID       uuid.UUID
	Author         Author
	Content        string
	Reactions      []Reaction
	CreatedAt      time.Time
	UpdatedAt      time.Time
	FailureMessage []string
}

type CommentOption func(*Comment)

func NewComment(opts ...CommentOption) *Comment {

	c := &Comment{}

	for _, opt := range opts {
		opt(c)
	}
	return c
}

func CommentWithID(id uuid.UUID) CommentOption {
	return func(c *Comment) {
		c.ID = id
	}
}

func CommentWithPostID(postID uuid.UUID) CommentOption {
	return func(c *Comment) {
		c.PostID = postID
	}
}

func CommentWithParentID(parentID uuid.UUID) CommentOption {
	return func(c *Comment) {
		c.ParentID = parentID
	}
}

func CommentWithAuthorID(authorID uuid.UUID) CommentOption {
	return func(c *Comment) {
		c.Author.ID = authorID
	}
}

func CommentWithAuthorName(authorName string) CommentOption {
	return func(c *Comment) {
		c.Author.Name = authorName
	}
}

func CommentWithAuthorImage(authorImage string) CommentOption {
	return func(c *Comment) {
		c.Author.Image = authorImage
	}
}

func CommentWithContent(content string) CommentOption {
	return func(c *Comment) {
		c.Content = content
	}
}

func CommentWithReactions(reactions []Reaction) CommentOption {
	return func(c *Comment) {
		c.Reactions = reactions
	}
}

func CommentWithCreatedAt(createdAt time.Time) CommentOption {
	return func(c *Comment) {
		c.CreatedAt = createdAt
	}
}

func CommentWithUpdatedAt(updatedAt time.Time) CommentOption {
	return func(c *Comment) {
		c.UpdatedAt = updatedAt
	}
}

func (c *Comment) Validate() {
	if c.Content == "" || len(c.Content) > 1000 {
		c.FailureMessage = append(c.FailureMessage, "Comment content must be between 1 and 1000 characters")
	}
}
