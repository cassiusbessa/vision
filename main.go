package main

import (
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
	r.POST("/posts", postController.CreatePost)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
