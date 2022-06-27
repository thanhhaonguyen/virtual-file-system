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

	// get Postgres DB connection
	db, err := models.GetDatabaseConnection()
	if util.HandleErrorInternalServer(c, err) {
		log.Printf("Failed to connect to database due to [Error]: %v", err)
		return
	}

	// query to get all files
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

	// get Postgres DB connection
	db, err := models.GetDatabaseConnection()
	if util.HandleErrorInternalServer(c, err) {
		log.Printf("Failed to connect to database due to [Error]: %v", err)
		return
	}

	// query to get file with fileId
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

	// parse request body into file model
	var file models.File
	err = json.Unmarshal(jsonData, &file)
	if util.HandleErrorBadRequest(c, err) {
		log.Printf("Failed to parse Json to file due to [Error]: %v", err)
		return
	}

	// get Postgres DB connection
	db, err := models.GetDatabaseConnection()
	if util.HandleErrorInternalServer(c, err) {
		log.Printf("Failed to connect to database due to [Error]: %v", err)
		return
	}

	// check if there are any file name already exists
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
	// check if there are any folder name already exists
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

	// query to create a new file
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

	// parse request body into file model
	var updateFile models.File
	err = json.Unmarshal(jsonData, &updateFile)
	if util.HandleErrorBadRequest(c, err) {
		log.Printf("Failed to parse Json to updateFile due to [Error]: %v", err)
		return
	}

	// get Postgres DB connection
	db, err := models.GetDatabaseConnection()
	if util.HandleErrorInternalServer(c, err) {
		log.Printf("Failed to connect to database due to [Error]: %v", err)
		return
	}

	// check if there is any record contain fileId
	var file models.File
	err = db.Where("id = ?", fileId).First(&file).Error
	if util.HandleErrorBadRequest(c, err) {
		log.Printf("No Record found due to [Error]: %v", err)
		return
	}

	// query to update file
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

	// get Postgres DB connection
	db, err := models.GetDatabaseConnection()
	if util.HandleErrorInternalServer(c, err) {
		log.Printf("Failed to connect to database due to [Error]: %v", err)
		return
	}

	// check if there is any record contain fileIds
	var file models.File
	err = db.Where("id = ?", fileId).First(&file).Error
	if util.HandleErrorBadRequest(c, err) {
		log.Printf("No Record found due to [Error]: %v", err)
		return
	}

	// query to delete file
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
