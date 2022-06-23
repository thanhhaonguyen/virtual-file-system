package v1

import (
	"encoding/json"
	"errors"
	"file-system-api/models"
	util "file-system-api/utils"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

func GetFolders(c *gin.Context) {
	var folders []models.Folder

	db, err := models.GetDatabaseConnection()
	if util.HandleErrorInternalServer(c, err) {
		log.Printf("Failed to connect to database due to [Error]: %v", err)
		return
	}

	err = db.Find(&folders).Error
	if util.HandleErrorBadRequest(c, err) {
		log.Printf("Failed to get folder list due to [Error]: %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success",
		"data":    folders,
	})
}

func GetFolder(c *gin.Context) {
	folderId := c.Param("id")

	var folder models.Folder

	db, err := models.GetDatabaseConnection()
	if util.HandleErrorInternalServer(c, err) {
		log.Printf("Failed to connect to database due to [Error]: %v", err)
		return
	}

	err = db.Where("id = ?", folderId).First(&folder).Error
	if util.HandleErrorBadRequest(c, err) {
		log.Printf("No Record found due to [Error]: %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success",
		"data":    folder,
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
		log.Printf("Failed to parse Json to folder due to [Error]: %v", err)
		return
	}

	db, err := models.GetDatabaseConnection()
	if util.HandleErrorInternalServer(c, err) {
		log.Printf("Failed to connect to database due to [Error]: %v", err)
		return
	}

	var folders []models.Folder
	folderName := folder.Name
	folderParentId := folder.ParentId

	err = db.Where("parent_id = ?", folderParentId).Find(&folders).Error
	if util.HandleErrorBadRequest(c, err) {
		log.Printf("No Record found due to [Error]: %v", err)
		return
	}

	for _, e := range folders {
		if folderName == e.Name {
			err = errors.New("the folder name already exists")
		}
		if util.HandleErrorBadRequest(c, err) {
			log.Printf("Failed to create new folder due to [Error]: %v", err)
			return
		}
	}

	err = db.Create(&folder).Error
	if util.HandleErrorBadRequest(c, err) {
		log.Printf("Failed to create new folder due to [Error]: %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success",
	})
}

func UpdateFolder(c *gin.Context) {
	folderId := c.Param("id")

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if util.HandleErrorBadRequest(c, err) {
		log.Printf("Failed to get Request Body due to [Error]: %v", err)
		return
	}

	updateFolder := models.Folder{}
	err = json.Unmarshal(jsonData, &updateFolder)
	if util.HandleErrorBadRequest(c, err) {
		log.Printf("Failed to parse Json to updateFolder due to [Error]: %v", err)
		return
	}

	db, err := models.GetDatabaseConnection()
	if util.HandleErrorInternalServer(c, err) {
		log.Printf("Failed to connect to database due to [Error]: %v", err)
		return
	}

	folder := models.Folder{}
	err = db.Where("id = ?", folderId).First(&folder).Error
	if util.HandleErrorBadRequest(c, err) {
		log.Printf("No Record found due to [Error]: %v", err)
		return
	}

	err = db.Model(&folder).Updates(&updateFolder).Error
	if util.HandleErrorBadRequest(c, err) {
		log.Printf("Failed to update folder due to [Error]: %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success",
	})
}

func DeleteFolder(c *gin.Context) {
	folderId := c.Param("id")

	db, err := models.GetDatabaseConnection()
	if util.HandleErrorInternalServer(c, err) {
		log.Printf("Failed to connect to database due to [Error]: %v", err)
		return
	}

	folder := models.Folder{}
	err = db.Where("id = ?", folderId).First(&folder).Error
	if util.HandleErrorBadRequest(c, err) {
		log.Printf("No Record found due to [Error]: %v", err)
		return
	}

	err = db.Delete(&folder).Error
	if util.HandleErrorBadRequest(c, err) {
		log.Printf("Failed to delete folder due to [Error]: %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success",
	})
}
