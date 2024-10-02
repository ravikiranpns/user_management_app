package main

import (
    "user_management_app/database"
    "user_management_app/models"
    "user_management_app/routes"   
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    // Connect to the database
    database.ConnectDB()
     // Connect to MongoDB
    database.ConnectMongoDB()
    // Create an admin user if it doesn't exist
    createAdminUser()

    // Setup routes
    routes.SetupRoutes(router)

    // Start the server
    router.Run(":8080")
}

// Function to create admin user
func createAdminUser() {
    var admin models.User
    if err := database.DB.Where("email = ?", "admin@example.com").First(&admin).Error; err != nil {
        admin = models.User{
            Name:     "Admin User",
            Email:    "admin@example.com",
            Password: "password123",  // Consider hashing this password in production
        }
        database.DB.Create(&admin)
        println("Admin user created")
    }
}
