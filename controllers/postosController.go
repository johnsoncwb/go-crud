package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/johnsoncwb/go-crud/initializers"
	"github.com/johnsoncwb/go-crud/models"
)

func PostsCreate(c *gin.Context) {

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)
	post := models.Post{Title: body.Title, Body: body.Title}

	result := initializers.DB.Create(&post)

	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"post": post,
	})
}

func PostsIndex(c *gin.Context) {
	var posts []models.Post

	status := initializers.DB.Find(&posts)

	if status.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"posts": posts,
	})

}

func PostsShow(c *gin.Context) {

	id := c.Param("id")

	var post models.Post

	result := initializers.DB.First(&post, id)

	if result.Error != nil {
		c.Status(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})

}

func PostsUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		Body  string
		Title string
	}

	c.Bind(&body)

	var post models.Post
	firstResult := initializers.DB.First(&post, id)

	if firstResult.Error != nil {
		c.Status(http.StatusNotFound)
		return
	}

	updateResult := initializers.DB.Model(&post).Updates(models.Post{
		Title: body.Title,
		Body:  body.Body,
	})

	if updateResult.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"post": post,
	})

}

func PostDelete(c *gin.Context) {
	id := c.Param("id")

	result := initializers.DB.Delete(&models.Post{}, id)

	if result.Error != nil {
		c.Status(http.StatusBadRequest)
		return
	}

	c.Status(http.StatusOK)
}

func Alive(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": "alive",
	})
}
