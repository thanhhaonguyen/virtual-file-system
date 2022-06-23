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

func GetFiles(c *gin.Context) {
	var files []models.File

	db, err := models.GetDatabaseConnection()
	if util.HandleErrorInternalServer(c, err) {
		log.Printf("Failed to connect to database due to [Error]: %v", err)
		return
	}

	err = db.Find(&files).Error
	if util.HandleErrorBadRequest(c, err) {
		log.Printf("Failed to get file list due to [Error]: %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success",
		"data":    files,
	})
}

func GetFile(c *gin.Context) {
	fileId := c.Param("id")

	var file models.File

	db, err := models.GetDatabaseConnection()
	if util.HandleErrorInternalServer(c, err) {
		log.Printf("Failed to connect to database due to [Error]: %v", err)
		return
	}

	err = db.Where("id = ?", fileId).First(&file).Error
	if util.HandleErrorBadRequest(c, err) {
		log.Printf("No Record found due to [Error]: %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success",
		"data":    file,
	})
}

func CreateFile(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if util.HandleErrorBadRequest(c, err) {
		log.Printf("Failed to get Request Body due to [Error]: %v", err)
		return
	}

	var file models.File
	err = json.Unmarshal(jsonData, &file)
	if util.HandleErrorBadRequest(c, err) {
		log.Printf("Failed to parse Json to file due to [Error]: %v", err)
		return
	}

	db, err := models.GetDatabaseConnection()
	if util.HandleErrorInternalServer(c, err) {
		log.Printf("Failed to connect to database due to [Error]: %v", err)
		return
	}

	var files []models.File
	fileName := file.Name
	fileFolderId := file.FolderId
	db.Where("folder_id = ?", fileFolderId).Find(&files)
	if len(files) > 0 {
		for _, e := range files {
			if fileName == e.Name {
				err = errors.New("the file name already exists")
			}
			if util.HandleErrorBadRequest(c, err) {
				log.Printf("Failed to create new file due to [Error]: %v", err)
				return
			}
		}
	}
	var folders []models.Folder
	db.Where("parent_id = ?", fileFolderId).Find(&folders)
	if len(folders) > 0 {
		for _, e := range folders {
			if fileName == e.Name {
				err = errors.New("the file name already exists")
			}
			if util.HandleErrorBadRequest(c, err) {
				log.Printf("Failed to create new file due to [Error]: %v", err)
				return
			}
		}
	}

	err = db.Create(&file).Error
	if util.HandleErrorBadRequest(c, err) {
		log.Printf("Failed to create new file due to [Error]: %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success",
	})
}

func UpdateFile(c *gin.Context) {
	fileId := c.Param("id")

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if util.HandleErrorBadRequest(c, err) {
		log.Printf("Failed to get Request Body due to [Error]: %v", err)
		return
	}

	var updateFile models.File
	err = json.Unmarshal(jsonData, &updateFile)
	if util.HandleErrorBadRequest(c, err) {
		log.Printf("Failed to parse Json to updateFile due to [Error]: %v", err)
		return
	}

	db, err := models.GetDatabaseConnection()
	if util.HandleErrorInternalServer(c, err) {
		log.Printf("Failed to connect to database due to [Error]: %v", err)
		return
	}

	var file models.File
	err = db.Where("id = ?", fileId).First(&file).Error
	if util.HandleErrorBadRequest(c, err) {
		log.Printf("No Record found due to [Error]: %v", err)
		return
	}

	err = db.Model(&file).Updates(&updateFile).Error
	if util.HandleErrorBadRequest(c, err) {
		log.Printf("Failed to update file due to [Error]: %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success",
	})
}

func DeleteFile(c *gin.Context) {
	fileId := c.Param("id")

	db, err := models.GetDatabaseConnection()
	if util.HandleErrorInternalServer(c, err) {
		log.Printf("Failed to connect to database due to [Error]: %v", err)
		return
	}

	var file models.File
	err = db.Where("id = ?", fileId).First(&file).Error
	if util.HandleErrorBadRequest(c, err) {
		log.Printf("No Record found due to [Error]: %v", err)
		return
	}

	err = db.Delete(&file).Error
	if util.HandleErrorBadRequest(c, err) {
		log.Printf("Failed to delete file due to [Error]: %v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Success",
	})
}
