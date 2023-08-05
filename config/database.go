package config

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Database struct {
	Conn   *sql.DB
	Logger zerolog.Logger
}

type Config struct {
	Host     string
	Port     int
	Username string
	Password string
	DbName   string
	Logger   zerolog.Logger
}

func DatabaseConnection() *gorm.DB {

	// Initialize DB
	err := godotenv.Load()
	if err != nil {
		log.Info().Msg("Error loading .env file")
	}

	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", os.Getenv("DB_HOST"), 5432, os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"))
	db, err := gorm.Open(postgres.Open(sqlInfo), &gorm.Config{})
	// helper.ErrorPanic(err)

	return db
}

func EnvCloudName() string {
	err := godotenv.Load()
	if err != nil {
		log.Info().Msg("Error loading .env file")
	}
	return os.Getenv("CLOUDINARY_CLOUD_NAME")
}

func EnvCloudAPIKey() string {
	err := godotenv.Load()
	if err != nil {
		log.Info().Msg("Error loading .env file")
	}
	return os.Getenv("CLOUDINARY_API_KEY")
}

func EnvCloudAPISecret() string {
	err := godotenv.Load()
	if err != nil {
		log.Info().Msg("Error loading .env file")
	}
	return os.Getenv("CLOUDINARY_API_SECRET")
}

func EnvCloudUploadFolder() string {
	err := godotenv.Load()
	if err != nil {
		log.Info().Msg("Error loading .env file")
	}
	return os.Getenv("CLOUDINARY_UPLOAD_FOLDER")
}
