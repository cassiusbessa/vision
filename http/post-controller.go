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

func (controller *PostController) GetPosts(c *gin.Context) {
	posts, err := controller.postService.LoadOrderedPosts(
		&dto.LoadOrderedPostsQuery{
			Limit:  10,
			Offset: 0,
		},
	)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (controller *PostController) ReactToPost(c *gin.Context) {
	var command dto.ReactToPostCommand
	if err := c.ShouldBindJSON(&command); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := controller.postService.ReactToPost(&command)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (controller *PostController) RemovePostReaction(c *gin.Context) {
	var command dto.RemovePostReactionCommand
	if err := c.ShouldBindJSON(&command); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := controller.postService.RemovePostReaction(&command)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response)
}
