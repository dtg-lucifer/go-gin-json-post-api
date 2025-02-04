package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GetAllPosts(c *gin.Context, db *gorm.DB) {

}

func GetPostById(c *gin.Context, db *gorm.DB) {
  _ = c.Param("id")
}

func CreatePost(c *gin.Context, db *gorm.DB) {}

func DeletePost(c *gin.Context, db *gorm.DB) {}
