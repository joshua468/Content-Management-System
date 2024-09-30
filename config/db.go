package config

import (
    "fmt"
    "os"

    "github.com/joho/godotenv"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "github.com/joshua468/content-management-system/models"
)

var DB *gorm.DB

func InitDB() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if err != nil {
        fmt.Println("Error loading .env file")
    }

    dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
        os.Getenv("DB_PORT"),
    )

    DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        fmt.Println("Error connecting to the database: ", err)
    }

    DB.AutoMigrate(&models.User{}, &models.Post{})
}
