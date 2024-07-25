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

	response, err := controller.postService.CreatePost(&command)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (controller *PostController) UpdatePost(c *gin.Context) {
	var command dto.UpdatePostCommand
	if err := c.ShouldBindJSON(&command); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := controller.postService.UpdatePost(&command)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response)
}
