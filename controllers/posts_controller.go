package controllers

import (
	"fmt"
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
  id, err := strconv.Atoi(c.Param("id"))
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "message": "Error parsing the id, make sure you provided that as an integer",
    })
    return
  }

  var post models.Post
  db.First(&post, id)

  if post.ID == 0 {
    c.JSON(http.StatusNotFound, gin.H{
      "message": fmt.Sprintf("Post with the id: %d, not found in the DB", id),
    })
    return
  }

  c.JSON(http.StatusFound, gin.H{
    "message": "Found!",
    "post": post,
  })
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

func DeletePost(c *gin.Context, db *gorm.DB) {
  id, err := strconv.Atoi(c.Param("id"))
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "message": "Make sure you provide an id",
    })
    return
  }

  db.Delete(&models.Post{}, id)

  c.JSON(http.StatusOK, gin.H{
    "message": "Deleted the post!",
  })
}

func UpdatePost(c *gin.Context, db *gorm.DB) {
  var rBody models.CreatePostDTO
  if err := c.BindJSON(&rBody); err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "message": "Cannot parse the body, make sure you passed the right data",
    })
    return
  }

  id, err := strconv.Atoi(c.Param("id")) 
  if err != nil {
    c.JSON(http.StatusBadRequest, gin.H{
      "message": "Make sure you provide an id",
    })
    return
  }

  var post models.Post
  db.First(&post, id)

  if post.ID == 0 {
    c.JSON(http.StatusNotFound, gin.H{
      "message": fmt.Sprintf("Post with the id: %d, not found in the DB", id),
    })
    return
  }

  db.Model(&post).Updates(models.Post{
    Title: rBody.Title,
    Body: rBody.Body,
  })

  c.JSON(http.StatusOK, gin.H{
    "message": "Updated the post!",
    "data": post,
  })
}
