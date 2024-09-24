package dtos

type AuthorLoadedResponse struct {
	AuthorID    string `json:"author_id"`
	AuthorName  string `json:"author_name"`
	AuthorImage string `json:"author_image"`
}
