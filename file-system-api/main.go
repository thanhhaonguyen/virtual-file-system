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
	r.Use(cors.MiddlewareCors())

	models.CreateDBConnection()

	// folder APIs
	r.GET("/folder", controller.GetFolders)
	r.GET("/folder/:id", controller.GetFolder)
	r.POST("/folder", controller.CreateFolder)
	r.PUT("/folder/:id", controller.UpdateFolder)
	r.DELETE("/folder/:id", controller.DeleteFolder)

	r.Run(initDomain())
}
