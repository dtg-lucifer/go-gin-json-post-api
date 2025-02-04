package main

import (
	"strings"

	"github.com/dtg-lucifer/go-gorm-gin-json-api/routes"
	"github.com/dtg-lucifer/go-gorm-gin-json-api/storage"
	"github.com/dtg-lucifer/go-gorm-gin-json-api/utils"

	"github.com/gin-gonic/gin"
)

func main() {
  utils.LoadEnvVariables()
  store := storage.NewStorage()

  engine := gin.Default()

  r := routes.SetupRoutes(store, engine)
  port:= utils.Must(utils.GetEnv("PORT"))

  r.Engine.Run(strings.Join([]string{":", port}, ""))
}
