package entities

import (
	"github.com/google/uuid"
)

type ReactionType string

const (
	Like    ReactionType = "Like"
	Dislike ReactionType = "Dislike"
	Love    ReactionType = "Love"
	Wow     ReactionType = "Wow"
	Angry   ReactionType = "Angry"
)

type Reaction struct {
	ID             uuid.UUID
	ProjectID      uuid.UUID
	UserID         uuid.UUID
	Type           ReactionType
	FailureMessage []string
}

func NewReaction(projectID, userID uuid.UUID, reactionType ReactionType) *Reaction {
	return &Reaction{
		ID:             uuid.New(),
		ProjectID:      projectID,
		UserID:         userID,
		Type:           reactionType,
		FailureMessage: []string{},
	}
}

func (r *Reaction) Validate() {
	if r.Type != Like && r.Type != Dislike && r.Type != Love && r.Type != Wow && r.Type != Angry {
		r.FailureMessage = append(r.FailureMessage, "Invalid reaction type")
	}
}
