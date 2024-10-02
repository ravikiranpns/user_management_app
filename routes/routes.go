package routes

import (
    "github.com/gin-gonic/gin"
    "user_management_app/controllers"
)

func SetupRoutes(router *gin.Engine) {
    // Public route for admin login
    router.POST("/admin/login", controllers.AdminLogin)

    // User-related routes
    router.GET("/users", controllers.GetUsers)
    router.POST("/users", controllers.CreateUser)
    router.PUT("/users/:id", controllers.UpdateUser)
    router.DELETE("/users/:id", controllers.DeleteUser)
}
