package controllers

import (
	"github/merkuluf/go-crud/initializers"
	"github/merkuluf/go-crud/models"

	"github.com/gin-gonic/gin"
)

func PostsCreate(c *gin.Context) {

	// get data of req body
	var body struct {
		Title string
		Body  string
	}

	c.Bind(&body)

	// create a post
	post := models.Post{Title: body.Title, Body: body.Body}
	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// return it
	c.JSON(200, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	// get the posts
	var posts []models.Post
	initializers.DB.Find(&posts)

	//respond with it
	c.JSON(200, gin.H{
		"posts": posts,
	})
}

func PostsShow(c *gin.Context) {
	// get id off url
	id := c.Param("id")

	// get the post
	var post models.Post
	initializers.DB.First(&post, id)

	c.JSON(200, gin.H{
		"post": post,
	})

}
