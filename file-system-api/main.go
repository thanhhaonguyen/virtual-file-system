package main

import (
	controller "file-system-api/controllers/v1"
	"file-system-api/middlewares/cors"
	"file-system-api/models"
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/gin-gonic/gin"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("no env option")
	}
}

func initDomain() string {
	appEnv := os.Getenv("APP_ENV")
	appPort := os.Getenv("PORT")

	// for local
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

	// cors middleware
	r.Use(cors.MiddlewareCors())

	// connect to Postgres DB
	models.CreateDBConnection()

	// folder APIs
	r.GET("/folder", controller.GetFolders)
	r.GET("/folder-by-parent/:id", controller.GetFoldersByParentId)
	r.GET("/folder/:id", controller.GetFolder)
	r.POST("/folder", controller.CreateFolder)
	r.PUT("/folder/:id", controller.UpdateFolder)
	r.DELETE("/folder/:id", controller.DeleteFolder)

	// file APIs
	r.GET("/file", controller.GetFiles)
	r.GET("/file/:id", controller.GetFile)
	r.POST("/file", controller.CreateFile)
	r.PUT("/file/:id", controller.UpdateFile)
	r.DELETE("/file/:id", controller.DeleteFile)

	r.Run(initDomain())
}
