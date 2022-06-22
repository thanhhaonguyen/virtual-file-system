package main

import (
	controller "file-system-api/controllers/v1"
	"file-system-api/middlewares/cors"
	"file-system-api/models"
	"github.com/joho/godotenv"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("no env option")
	}
}

func initDomain() string {
	appEnv := os.Getenv("APP_ENV")
	appPort := os.Getenv("API_PORT")

	domain := "localhost"

	if appEnv == "prod" {
		gin.SetMode(gin.ReleaseMode)
		domain = ""
	}
	if appEnv == "dev" {
		domain = ""
	}
	return domain + ":" + appPort
}

func main() {
	r := gin.Default()
	r.Use(cors.CorsMiddleware())

	models.CreateDBConnection()

	r.GET("/folders", controller.GetFolderList)
	r.POST("/folder", controller.CreateFolder)

	r.Run(initDomain())
}
