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
	PostID         uuid.UUID
	ParentID       uuid.NullUUID
	UserID         uuid.UUID
	Type           ReactionType
	FailureMessage []string
}

type ReactionOption func(*Reaction)

func NewReaction(opts ...ReactionOption) *Reaction {

	r := &Reaction{}

	for _, opt := range opts {
		opt(r)
	}
	return r
}

func ReactionWithID(id uuid.UUID) ReactionOption {
	return func(r *Reaction) {
		r.ID = id
	}
}

func ReactionWithPostID(postID uuid.UUID) ReactionOption {
	return func(r *Reaction) {
		r.PostID = postID
	}
}

func ReactionWithParentID(parentID uuid.NullUUID) ReactionOption {
	return func(r *Reaction) {
		r.ParentID = parentID
	}
}

func ReactionWithUserID(userID uuid.UUID) ReactionOption {
	return func(r *Reaction) {
		r.UserID = userID
	}
}

func ReactionWithReactionType(reactionType ReactionType) ReactionOption {
	return func(r *Reaction) {
		r.Type = reactionType
	}
}

func (r *Reaction) Validate() {
	if r.Type != Like && r.Type != Dislike && r.Type != Love && r.Type != Wow && r.Type != Angry {
		r.FailureMessage = append(r.FailureMessage, "Invalid reaction type")
	}
}
