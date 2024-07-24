package http

import (
	"net/http"

	dto "github.com/cassiusbessa/vision-social-media/domain/service/dtos"
	ports "github.com/cassiusbessa/vision-social-media/domain/service/ports/input"
	"github.com/gin-gonic/gin"
)

type PostController struct {
	postService ports.PostService
}

func NewPostController(postService ports.PostService) *PostController {
	return &PostController{
		postService: postService,
	}
}

func (controller *PostController) CreatePost(c *gin.Context) {
	var command dto.CreatePostCommand
	if err := c.ShouldBindJSON(&command); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := controller.postService.CreatePost(&command)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}
