package main

import (
	"github.com/dtg-lucifer/go-gorm-gin-json-api/models"
	"github.com/dtg-lucifer/go-gorm-gin-json-api/storage"
	"github.com/dtg-lucifer/go-gorm-gin-json-api/utils"
	"gorm.io/gorm"
)

func MigrateModels(db *gorm.DB) {
  utils.Must[any](nil, db.AutoMigrate(&models.Post{}))
}

func main() {
  db := storage.InitDb()

  utils.Must[any](nil, db.AutoMigrate(&models.Post{}))
}
