package service

import (
	outputPorts "github.com/cassiusbessa/vision-social-media/domain/service/ports/output"
)

type PostService struct {
	postRepo outputPorts.PostRepository
}

func NewPostService(postRepo outputPorts.PostRepository) *PostService {
	return &PostService{
		postRepo: postRepo,
	}
}
