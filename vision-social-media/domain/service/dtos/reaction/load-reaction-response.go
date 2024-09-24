package dtos

import "github.com/cassiusbessa/vision-social-media/domain/core/entities"

type LoadReactionResponse struct {
	ID        string                `json:"id"`
	PostID    string                `json:"post_id"`
	ParentID  string                `json:"parent_id"`
	UserID    string                `json:"user_id"`
	Type      entities.ReactionType `json:"type"`
	CreatedAt string                `json:"created_at"`
}
