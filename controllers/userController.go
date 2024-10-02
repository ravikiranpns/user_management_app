package controllers

import (
    "context"
    "github.com/gin-gonic/gin"
    "go.mongodb.org/mongo-driver/bson"
    "net/http"
    "time"
    "user_management_app/database"
    "user_management_app/models"
    "github.com/dgrijalva/jwt-go"
)

var jwtKey = []byte("your_secret_key")

// Claims struct to use for JWT token
type Claims struct {
    Email string `json:"email"`
    jwt.StandardClaims
}

// CreateUser - Create a new user in the database and log the event
func CreateUser(c *gin.Context) {
    var user models.User
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    database.DB.Create(&user)
    c.JSON(http.StatusOK, user)

    // Asynchronous logging of the user creation
    go func() {
        logCollection := database.MongoClient.Database("user_logs").Collection("logs")
        logEntry := bson.M{
            "user_id":   user.ID,
            "event":     "User Created",
            "data":      user,
            "timestamp": time.Now(),
        }
        logCollection.InsertOne(context.Background(), logEntry)
    }()
}

// GetUsers - Fetch all users and log the event
func GetUsers(c *gin.Context) {
    var users []models.User
    database.DB.Find(&users)
    c.JSON(http.StatusOK, users)

    // Asynchronous logging of fetching users
    go func() {
        logCollection := database.MongoClient.Database("user_logs").Collection("logs")
        logEntry := bson.M{
            "event":     "Fetched All Users",
            "timestamp": time.Now(),
        }
        logCollection.InsertOne(context.Background(), logEntry)
    }()
}

// UpdateUser - Update user data and log the event
func UpdateUser(c *gin.Context) {
    var user models.User
    id := c.Param("id")

    if err := database.DB.First(&user, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    database.DB.Save(&user)
    c.JSON(http.StatusOK, user)

    // Asynchronous logging of the user update
    go func() {
        logCollection := database.MongoClient.Database("user_logs").Collection("logs")
        logEntry := bson.M{
            "user_id":   user.ID,
            "event":     "User Updated",
            "data":      user,
            "timestamp": time.Now(),
        }
        logCollection.InsertOne(context.Background(), logEntry)
    }()
}

// DeleteUser - Delete a user and log the event
func DeleteUser(c *gin.Context) {
    var user models.User
    id := c.Param("id")

    if err := database.DB.First(&user, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
        return
    }

    database.DB.Delete(&user)
    c.JSON(http.StatusOK, gin.H{"message": "User deleted"})

    // Asynchronous logging of the user deletion
    go func() {
        logCollection := database.MongoClient.Database("user_logs").Collection("logs")
        logEntry := bson.M{
            "user_id":   user.ID,
            "event":     "User Deleted",
            "timestamp": time.Now(),
        }
        logCollection.InsertOne(context.Background(), logEntry)
    }()
}

// AdminLogin - Authenticate admin and return JWT token

func AdminLogin(c *gin.Context) {
    var creds struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    if err := c.ShouldBindJSON(&creds); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    // Validate credentials against the database
    var admin models.User
    if err := database.DB.Where("email = ?", creds.Email).First(&admin).Error; err != nil || admin.Password != creds.Password {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
        return
    }

    // Return success response for now
    c.JSON(http.StatusOK, gin.H{"message": "Login successful!"})
}
