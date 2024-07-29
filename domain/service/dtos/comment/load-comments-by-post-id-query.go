package dtos

type LoadOrderedCommentsQuery struct {
	Limit  int32
	Offset int32
	PostID string
}
