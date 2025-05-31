package main

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthMiddleware(jwtKey []byte) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header required"})
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer {token}"})
			c.Abort()
			return
		}

		tokenString := parts[1]
		claims := &jwt.RegisteredClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

		if err != nil {
			switch {
			case errors.Is(err, jwt.ErrTokenSignatureInvalid):
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token signature"})
			case errors.Is(err, jwt.ErrTokenMalformed):
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Malformed token"})
			case errors.Is(err, jwt.ErrTokenExpired), errors.Is(err, jwt.ErrTokenNotValidYet):
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token is expired or not active yet"})
			default:
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Couldn't handle this token: " + err.Error()})
			}
			c.Abort()
			return
		}

		if !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Set values from the standard claims
		c.Set("userID", claims.ID)
		c.Set("userEmail", claims.Subject)

		c.Next()
	}
}
