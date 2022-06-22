package v1

import (
	"encoding/json"
	"file-system-api/models"
	util "file-system-api/utils"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

func GetFolderList(c *gin.Context) {
	var folders []models.Folder

	db, err := models.GetDatabaseConnection()
	if util.HandleErrorInternalServer(c, err) {
		log.Printf("Failed to connect to database due to [Error]: %v", err)
		return
	}

	db.Find(&folders)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success",
		"data":    folders,
	})
}

func CreateFolder(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if util.HandleErrorBadRequest(c, err) {
		log.Printf("Failed to get Request Body due to [Error]: %v", err)
		return
	}

	folder := models.Folder{}
	err = json.Unmarshal(jsonData, &folder)
	if util.HandleErrorBadRequest(c, err) {
		log.Printf("Failed to parse Json to createFolderDTO due to [Error]: %v", err)
		return
	}

	db, err := models.GetDatabaseConnection()
	if util.HandleErrorInternalServer(c, err) {
		log.Printf("Failed to connect to database due to [Error]: %v", err)
		return
	}

	db.Create(&folder)

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success",
	})
}
