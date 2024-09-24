package main

import (
	"log"

	httpNet "net/http"

	data "github.com/cassiusbessa/vision-social-media/data-access"
	sqlc "github.com/cassiusbessa/vision-social-media/data-access/sqlc-config"
	service "github.com/cassiusbessa/vision-social-media/domain/service/implementation"
	http "github.com/cassiusbessa/vision-social-media/http"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/consul/api"
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

	r.GET("/health", func(c *gin.Context) {
		c.JSON(httpNet.StatusOK, gin.H{
			"status": "healthy",
		})
	})

	// Registrar serviço no Consul
	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		log.Fatalf("Erro ao criar cliente Consul: %v", err)
	}

	registration := new(api.AgentServiceRegistration)
	registration.ID = "vision-social-media"
	registration.Name = "vision-social-media"
	registration.Port = 8888
	registration.Tags = []string{"social-media", "post-service"}
	registration.Address = "127.0.0.1"

	check := new(api.AgentServiceCheck)
	check.HTTP = "http://127.0.0.1:8888/health"
	check.Interval = "10s"
	check.Timeout = "1s"
	registration.Check = check

	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		log.Fatalf("Erro ao registrar serviço no Consul: %v", err)
	}

	log.Println("Server started on port 8888")
	if err := r.Run(":8888"); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
