package storage

import (
	"fmt"

	"github.com/dtg-lucifer/go-gorm-gin-json-api/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	DB *gorm.DB
}

func InitDb() *gorm.DB {

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		utils.Must(utils.GetEnv("DB_HOST")),
		utils.Must(utils.GetEnv("DB_USER")),
		utils.Must(utils.GetEnv("DB_PASS")),
		utils.Must(utils.GetEnv("DB_NAME")),
		utils.Must(utils.GetEnv("DB_PORT")),
		utils.Must(utils.GetEnv("DB_SSLMODE")),
	)

	db := utils.Must(gorm.Open(postgres.Open(dsn), &gorm.Config{}))

	return db
}

func NewStorage() *Storage {
	s := &Storage{
		DB: InitDb(),
	}

	return s
}
