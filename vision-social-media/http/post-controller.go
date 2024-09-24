package http

import (
	"net/http"
	"strconv"

	commentDTO "github.com/cassiusbessa/vision-social-media/domain/service/dtos/comment"
	postDTO "github.com/cassiusbessa/vision-social-media/domain/service/dtos/post"
	reactionDTO "github.com/cassiusbessa/vision-social-media/domain/service/dtos/reaction"
	ports "github.com/cassiusbessa/vision-social-media/domain/service/ports/input"
	"github.com/gin-gonic/gin"
)

type PostController struct {
	postService  ports.PostService
	tokenService ports.TokenService
}

func NewPostController(postService ports.PostService, tokenService ports.TokenService) *PostController {
	return &PostController{
		postService:  postService,
		tokenService: tokenService,
	}
}

func (controller *PostController) CreatePost(c *gin.Context) {
	var command postDTO.CreatePostCommand
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

func (controller *PostController) RemovePost(c *gin.Context) {
	var command postDTO.RemovePostCommand
	if err := c.ShouldBindJSON(&command); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := controller.postService.DeletePost(&command)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (controller *PostController) UpdatePost(c *gin.Context) {
	var command postDTO.UpdatePostCommand
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
	limitQuery := c.DefaultQuery("limit", "10")
	offsetQuery := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitQuery)
	if err != nil {
		limit = 10
	}
	offset, err := strconv.Atoi(offsetQuery)
	if err != nil {
		offset = 0
	}

	posts, err := controller.postService.LoadOrderedPosts(
		&postDTO.LoadOrderedPostsQuery{
			Limit:  int32(limit),
			Offset: int32(offset),
		},
	)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, posts)
}

func (controller *PostController) ReactToPost(c *gin.Context) {
	var command reactionDTO.ReactToPostCommand
	if err := c.ShouldBindJSON(&command); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := controller.tokenService.GetPayload(c.GetHeader("Authorization"))
	if err != nil {
		c.Error(err)
		return
	}
	command.UserID = userID

	response, err := controller.postService.ReactToPost(&command)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (controller *PostController) RemovePostReaction(c *gin.Context) {
	var command reactionDTO.RemovePostReactionCommand
	if err := c.ShouldBindJSON(&command); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := controller.tokenService.GetPayload(c.GetHeader("Authorization"))
	if err != nil {
		c.Error(err)
		return
	}
	command.UserID = userID

	response, err := controller.postService.RemovePostReaction(&command)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (controller *PostController) LoadPostReactionsByPostID(c *gin.Context) {
	postID := c.Param("postID")
	limitQuery := c.DefaultQuery("limit", "10")
	offsetQuery := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitQuery)
	if err != nil {
		limit = 10
	}
	offset, err := strconv.Atoi(offsetQuery)
	if err != nil {
		offset = 0
	}
	reactions, err := controller.postService.LoadPostReactionsByPostID(
		&reactionDTO.LoadOrderedReactionsQuery{
			PostID: postID,
			Limit:  int32(limit),
			Offset: int32(offset),
		},
	)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, reactions)
}

func (controller *PostController) AddCommentToPost(c *gin.Context) {
	var command commentDTO.AddCommentToPostCommand
	if err := c.ShouldBindJSON(&command); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := controller.tokenService.GetPayload(c.GetHeader("Authorization"))
	if err != nil {
		c.Error(err)
		return
	}
	command.AuthorID = userID

	response, err := controller.postService.AddCommentToPost(&command)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (controller *PostController) RemovePostComment(c *gin.Context) {
	var command commentDTO.RemovePostCommentCommand
	if err := c.ShouldBindJSON(&command); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := controller.tokenService.GetPayload(c.GetHeader("Authorization"))
	if err != nil {
		c.Error(err)
		return
	}
	command.UserID = userID

	response, err := controller.postService.RemovePostComment(&command)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, response)
}

func (controller *PostController) LoadPostCommentsByPostID(c *gin.Context) {
	postID := c.Param("postID")
	limitQuery := c.DefaultQuery("limit", "10")
	offsetQuery := c.DefaultQuery("offset", "0")

	limit, err := strconv.Atoi(limitQuery)
	if err != nil {
		limit = 10
	}
	offset, err := strconv.Atoi(offsetQuery)
	if err != nil {
		offset = 0
	}

	comments, err := controller.postService.LoadPostCommentsByPostID(
		&commentDTO.LoadOrderedCommentsQuery{
			PostID: postID,
			Limit:  int32(limit),
			Offset: int32(offset),
		},
	)
	if err != nil {
		c.Error(err)
		return
	}

	c.JSON(http.StatusOK, comments)
}
