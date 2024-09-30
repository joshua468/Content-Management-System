package main

import (
    "log"
    "github.com/gin-gonic/gin"
    "github.com/joshua468/content-management-system/config"
    "github.com/joshua468/content-management-system/routes"
    "github.com/joho/godotenv"
)

func main() {
    // Load the .env file
    if err := godotenv.Load(); err != nil {
        log.Fatal("Error loading .env file")
    }

    // Initialize the database
    config.InitDB()

    // Set up the Gin router
    r := gin.Default()
    routes.SetupRoutes(r)

    // Run the server
    if err := r.Run(":8080"); err != nil {
        log.Fatal("Server run failed:", err)
    }
}
