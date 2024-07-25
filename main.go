package main

import (
	"log"

	data "github.com/cassiusbessa/vision-social-media/data-access"
	sqlc "github.com/cassiusbessa/vision-social-media/data-access/sqlc-config"
	service "github.com/cassiusbessa/vision-social-media/domain/service"
	http "github.com/cassiusbessa/vision-social-media/http"
)

func main() {
	postgresDb := data.NewDbConn()
	queries := sqlc.New(postgresDb)
	postRepo := data.NewPostRepository(queries)

	postService := service.NewPostService(postRepo)

	postController := http.NewPostController(postService)

	sqlc.CreateTable(postgresDb)

	r := http.Router()

	r.Use(http.ErrorHandler())
	r.POST("/posts", postController.CreatePost)
	r.PUT("/posts", postController.UpdatePost)

	go func() {
		log.Println("Server started on port 8888")
	}()

	if err := r.Run(":8888"); err != nil {
		panic(err)
	}
}
