package database

import (
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() (*gorm.DB, error) {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	password := os.Getenv("PASSWORD")
	dbName := os.Getenv("DBNAME")

	config := fmt.Sprintf("host=%s user=%s password=%s database=%s sslmode=disable", host, user, password, dbName)

	gormDb, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		return nil, errors.New("database connection failure")
	}

	db, err := gormDb.DB()
	if err != nil {
		return nil, errors.New("database connection failure")
	}

	db.SetConnMaxIdleTime(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxLifetime(time.Hour)

	return gormDb, nil
}
