package dtos

type RemovePostReactionCommand struct {
	UserID     string `json:"user_id"`
	ReactionID string `json:"reaction_id"`
	PostID     string `json:"post_id"`
}
