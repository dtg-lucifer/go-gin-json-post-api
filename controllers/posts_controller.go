package controllers

import (
	"net/http"
	"strconv"

	"github.com/dtg-lucifer/go-gorm-gin-json-api/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllPosts(c *gin.Context, db *gorm.DB) {
  var posts []models.Post
  db.Find(&posts)

  if len(posts) < 1 {
    c.JSON(http.StatusOK, gin.H{
      "message": "There are no posts to show",
    })
    return
  }

  c.JSON(http.StatusOK, gin.H{
    "message": "Succesfully fetched the posts",
    "posts": posts,
  })
  return
}

func GetPostById(c *gin.Context, db *gorm.DB) {
  id := strconv.Atoi(c.Param("id"))
}

func CreatePost(c *gin.Context, db *gorm.DB) {
  var requestBody models.CreatePostDTO
  if err := c.BindJSON(&requestBody); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "message": "Unable to parse the body or the body in empty",
    })
    return
  }

  post := models.Post{ Title: requestBody.Title, Body: requestBody.Body }
  result := db.Create(&post)

  if result.Error != nil {
    c.JSON(http.StatusBadGateway, gin.H{
      "message": "Could not create the entry on database!",
    })
    return
  }

  c.JSON(http.StatusCreated, gin.H{
    "message": "Succesfully created the post!",
    "data": post,
  })
}

func DeletePost(c *gin.Context, db *gorm.DB) {}
