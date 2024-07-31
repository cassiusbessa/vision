package main

import (
	"log"

	data "github.com/cassiusbessa/vision-social-media/data-access"
	sqlc "github.com/cassiusbessa/vision-social-media/data-access/sqlc-config"
	service "github.com/cassiusbessa/vision-social-media/domain/service/implementation"
	http "github.com/cassiusbessa/vision-social-media/http"
)

func main() {
	postgresDb := data.NewDbConn()
	queries := sqlc.New(postgresDb)
	postRepo := data.NewPostRepository(queries, postgresDb)

	postService := service.NewPostService(postRepo)
	tokenService := service.NewTokenService()

	postController := http.NewPostController(postService, tokenService)

	sqlc.CreateTable(postgresDb)

	r := http.Router()

	r.Use(http.ErrorHandler())

	r.POST("/posts", postController.CreatePost)
	r.DELETE("/posts", postController.RemovePost)
	r.PUT("/posts", postController.UpdatePost)
	r.GET("/posts", postController.GetPosts)

	r.POST("/posts/react", postController.ReactToPost)
	r.DELETE("/posts/react", postController.RemovePostReaction)
	r.GET("/posts/:postID/react", postController.LoadPostReactionsByPostID)

	r.POST("/posts/comment", postController.AddCommentToPost)
	r.DELETE("/posts/comment", postController.RemovePostComment)
	r.GET("/posts/:postID/comment", postController.LoadPostCommentsByPostID)

	go func() {
		log.Println("Server started on port 8888")
	}()

	if err := r.Run(":8888"); err != nil {
		panic(err)
	}
}
