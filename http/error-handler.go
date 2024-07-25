package http

import (
	"log"
	"net/http"

	"github.com/cassiusbessa/vision-social-media/domain/service/errors"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		for _, err := range c.Errors {
			switch err.Err.(type) {
			case *errors.ValidationError:
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Err.Error()})
			case *errors.InvalidArgument:
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Err.Error()})
			case *errors.ResourceNotFound:
				c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": err.Err.Error()})
			case *errors.ResourceAlreadyExists:
				c.AbortWithStatusJSON(http.StatusConflict, gin.H{"error": err.Err.Error()})
			default:
				log.Println(err.Err)
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
			}
		}
	}
}
