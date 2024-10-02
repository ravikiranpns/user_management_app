package controllers

import (
    "net/http"
    "net/http/httptest"
    "strings"
    "testing"
    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "user_management_app/database"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"
    "user_management_app/models" 
)

func SetupTest() {
    // Initialize the test database connection
    dsn := "host=localhost user=user2 password=password2 dbname=test_user_management port=5432 sslmode=disable"
    var err error
    database.DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to test database")
    }

    // Migrate the User model for testing purposes
    database.DB.AutoMigrate(&models.User{})
}

// Test the CreateUser function
func TestCreateUser(t *testing.T) {
    SetupTest() // Set up the database for testing

    router := gin.Default()
    router.POST("/users", CreateUser)

    payload := `{"name": "Test User", "email": "test@example.com", "password": "password123"}`
    req, _ := http.NewRequest("POST", "/users", strings.NewReader(payload))
    req.Header.Set("Content-Type", "application/json")

    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
}

// Test the GetUsers function
func TestGetUsers(t *testing.T) {
    SetupTest() // Set up the database for testing

    router := gin.Default()
    router.GET("/users", GetUsers)

    req, _ := http.NewRequest("GET", "/users", nil)
    w := httptest.NewRecorder()
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
}
