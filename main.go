package main

import (
	"go-crud/m/controllers"
	"go-crud/m/intializers"

	"github.com/gin-gonic/gin"
)

func init() {
	intializers.LoadEnvVariables()
	intializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	// Create a new post
	r.POST("/posts", controllers.PostCreate)

	// List all posts
	r.GET("/posts", controllers.PostIndex)

	// Show a single post by ID
	r.GET("/posts/:id", controllers.PostsShow)

	// Update a post by ID
	r.PUT("/posts/:id", controllers.PostUpdate)

	// Delete a post by ID
	r.DELETE("/posts/:id", controllers.PostDelete)

	err := r.Run()
	if err != nil {
		return 
	} // listen and serve on 0.0.0.0:8080
}
