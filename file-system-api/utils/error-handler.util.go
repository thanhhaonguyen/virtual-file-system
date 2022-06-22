package utils

import (
	"github.com/gin-gonic/gin"
)

func HandleError(c *gin.Context, status int, message string, err error) bool {
	if err != nil {
		c.JSON(status, gin.H{
			"status":  status,
			"message": message,
			"errors":  err.Error(),
		})
		return true
	}
	return false
}

func HandleErrorBadRequest(c *gin.Context, err error) bool {
	return HandleError(c, 400, "Bad Request", err)
}

func HandleErrorUnauthorized(c *gin.Context, err error) bool {
	return HandleError(c, 401, "Unauthorized", err)
}

func HandleErrorInternalServer(c *gin.Context, err error) bool {
	return HandleError(c, 500, "Internal Server Error", err)
}
