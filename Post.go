package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Post struct {
	ID      int32  `json:"ID"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

var posts []Post

func PostInit() {
	posts = []Post{
		{ID: 1, Content: "This is post one", Author: "Shashank"},
		{ID: 2, Content: "This is post two", Author: "Someone else"},
	}
	fmt.Println(posts)
}

// ENDPOINT: /v1/posts/
func GetAllPost(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"posts": posts})
}

// ENDPOINT: /v1/posts/:id
func GetPostById(c *gin.Context) {
	postId64, err := strconv.ParseInt(c.Param("id"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": false, "message": err})
		return
	}
	postId := int32(postId64)

	var resultPost Post
	for _, post := range posts {
		if postId == post.ID {
			resultPost = post
		}
	}
	// If no result found
	if (Post{}) == resultPost {
		message := "No post found with id " + strconv.FormatInt(postId64, 10)
		c.JSON(http.StatusNoContent, gin.H{"status": false, "message": message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": true, "data": resultPost})
}
