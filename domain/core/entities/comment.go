package entities

import (
	"time"

	"github.com/google/uuid"
)

type Comment struct {
	ID             uuid.UUID
	ProjectID      uuid.UUID
	UserID         uuid.UUID
	Content        string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	FailureMessage []string
}

func NewComment(projectID, userID uuid.UUID, content string) *Comment {
	return &Comment{
		ID:             uuid.New(),
		ProjectID:      projectID,
		UserID:         userID,
		Content:        content,
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
		FailureMessage: []string{},
	}
}

func (c *Comment) Validate() {
	if c.Content == "" || len(c.Content) > 1000 {
		c.FailureMessage = append(c.FailureMessage, "Comment content cannot be empty")
	}
}
