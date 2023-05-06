package controllers

import (
	"github.com/gin-gonic/gin"
	"go-crud/m/intializers"
	"go-crud/m/models"
	"net/http"
)

func PostCreate(c *gin.Context) {

	var body struct {
		Body  string
		Title string
	}

	err := c.Bind(&body)
	if err != nil {
		return
	}

	post := models.Post{Title: body.Title, Body: body.Body}
	createResult := intializers.DB.Create(&post)

	if createResult.Error != nil {
		_ = c.AbortWithError(http.StatusBadRequest, createResult.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostIndex(c *gin.Context) {
	var posts []models.Post
	intializers.DB.Find(&posts)

	c.JSON(200, gin.H{
		"posts": posts,
	})

}

func PostsShow(c *gin.Context) {

	id := c.Param("id")

	var post models.Post
	intializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostUpdate(c *gin.Context) {

	id := c.Param("id")

	var body struct {
		Body  string
		Title string
	}

	err := c.Bind(&body)
	if err != nil {
		return
	}

	var post models.Post
	result := intializers.DB.First(&post, id)

	if result.Error != nil {
		_ = c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	post.Body = body.Body
	post.Title = body.Title
	updateResult := intializers.DB.Save(&post)

	if updateResult.Error != nil {
		_ = c.AbortWithError(http.StatusBadRequest, updateResult.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})
}

func PostDelete(c *gin.Context) {

	id := c.Param("id")

	var post models.Post
	result := intializers.DB.First(&post, id)

	if result.Error != nil {
		_ = c.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	deleteResult := intializers.DB.Delete(&post)

	if deleteResult.Error != nil {
		_ = c.AbortWithError(http.StatusBadRequest, deleteResult.Error)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Post deleted successfully",
	})
}
