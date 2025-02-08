package models

import "gorm.io/gorm"

type Post struct {
	gorm.Model
  Title string `binding:"required"`
  Body  string `binding:"required"`
}

type CreatePostDTO struct {
  Title string
  Body  string
}
