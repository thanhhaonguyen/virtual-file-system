package tests_test

import (
	controller "file-system-api/controllers/v1"
	model "file-system-api/models"
	"fmt"
	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
)

var dbConn *gorm.DB

var _ = Describe("[file-system-api.controller]", func() {
	// folder
	Describe("Test GetFolders", func() {
		Context("GetFolders with return no error", func() {
			It("Should return message with status 200", func() {
				// create gin context
				req, err := http.NewRequest("GET", "/folder", nil)
				Expect(err).To(BeNil())
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				c.Request = req

				// mock Postgres DB connection
				model.GetDatabaseConnection = func() (*gorm.DB, error) {
					return dbConn, nil
				}

				// test GetFolders
				controller.GetFolders(c)

				// assert data
				Expect(rr.Result().Status).To(Equal("200 OK"))
			})
		})
		Context("GetFolders with return error", func() {
			It("Should return message with status 500", func() {
				// create gin context
				req, err := http.NewRequest("GET", "/folder", nil)
				Expect(err).To(BeNil())
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				c.Request = req

				// mock Postgres DB connection
				model.GetDatabaseConnection = func() (*gorm.DB, error) {
					return dbConn, fmt.Errorf("failed to get DB connection")
				}

				// test GetFolders
				controller.GetFolders(c)

				// assert data
				Expect(rr.Result().Status).To(Equal("500 Internal Server Error"))
			})
		})
	})
	Describe("Test GetFoldersByParentId", func() {
		Context("GetFoldersByParentId with return no error", func() {
			It("Should return message with status 200", func() {
				// create gin context
				req, err := http.NewRequest("GET", "/folder-by-parent/0", nil)
				Expect(err).To(BeNil())
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				c.Request = req

				// mock Postgres DB connection
				model.GetDatabaseConnection = func() (*gorm.DB, error) {
					return dbConn, nil
				}

				// test GetFoldersByParentId
				controller.GetFoldersByParentId(c)

				// assert data
				Expect(rr.Result().Status).To(Equal("200 OK"))
			})
		})
		Context("GetFoldersByParentId with return error", func() {
			It("Should return message with status 500", func() {
				// create gin context
				req, err := http.NewRequest("GET", "/folder-by-parent/0", nil)
				Expect(err).To(BeNil())
				rr := httptest.NewRecorder()
				c, _ := gin.CreateTestContext(rr)
				c.Request = req

				// mock Postgres DB connection
				model.GetDatabaseConnection = func() (*gorm.DB, error) {
					return dbConn, fmt.Errorf("failed to get DB connection")
				}

				// test GetFoldersByParentId
				controller.GetFoldersByParentId(c)

				// assert data
				Expect(rr.Result().Status).To(Equal("500 Internal Server Error"))
			})
		})
	})
})
