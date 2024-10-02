package middleware

import (
    "github.com/gin-gonic/gin"
    "net/http"
    "github.com/dgrijalva/jwt-go"
    "strings"
    "time"
)

var jwtKey = []byte("your_secret_key")
// Claims struct to use for JWT token
type Claims struct {
    Email string `json:"email"`
    jwt.StandardClaims
}
// AuthMiddleware verifies the JWT token and allows access only to authenticated admins
func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
            c.Abort()
            return
        }

        tokenString := strings.Split(authHeader, " ")[1]
        claims := &Claims{}

        token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
            return jwtKey, nil
        })

        if err != nil || !token.Valid || claims.ExpiresAt < time.Now().Unix() {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
            c.Abort()
            return
        }

        // If everything is fine, proceed to the next handler
        c.Next()
    }
}
