package mappers

import (
	"github.com/cassiusbessa/vision-social-media/domain/core/entities"
	authorDTO "github.com/cassiusbessa/vision-social-media/domain/service/dtos/author"
)

func AuthorEnitityToLoadedResponse(author entities.Author) authorDTO.AuthorLoadedResponse {
	return authorDTO.AuthorLoadedResponse{
		AuthorID:    author.ID.String(),
		AuthorName:  author.Name,
		AuthorImage: author.Image,
	}
}
