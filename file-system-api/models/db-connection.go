package models

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
)

var dbConn *gorm.DB

func GetDSN() string {
	pgUser := os.Getenv("PG_USER")
	pgPassword := os.Getenv("PG_PASSWORD")
	pgDb := os.Getenv("PG_DB")
	pgHost := os.Getenv("PG_HOST")
	pgPort := os.Getenv("PG_PORT")
	pgSsl := os.Getenv("PG_SSL")
	pgTimeZone := os.Getenv("PG_TIMEZONE")

	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s", pgHost, pgUser, pgPassword, pgDb, pgPort, pgSsl, pgTimeZone)
}

func CreateDBConnection() {
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  GetDSN(),
		PreferSimpleProtocol: true,
	}), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   os.Getenv("PG_SCHEMA") + ".",
			SingularTable: true,
		}})
	if err != nil {
		panic("Error occurred while connecting with the database")
	}

	err = db.AutoMigrate(&Folder{}, &File{})
	if err != nil {
		panic("Error occurred while migrating schemas into the database")
	}

	dbConn = db
}

var GetDatabaseConnection = func() (*gorm.DB, error) {
	sqlDB, err := dbConn.DB()
	if err != nil {
		return dbConn, err
	}

	err = sqlDB.Ping()
	if err != nil {
		return dbConn, err
	}

	return dbConn, nil
}
