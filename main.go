package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	PostInit()

	registerRoute(r)

	// Running on port 8000
	r.Run(":8000")
}

func pingPong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "You're server is running"})
}

func registerRoute(r *gin.Engine) {
	// Home route handing
	r.GET("/", pingPong)
	postV1 := r.Group("/v1/posts/")
	{
		postV1.GET("/", GetAllPost)
		postV1.GET("/:id", GetPostById)
	}
}
