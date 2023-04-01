package middleware

import (
	"log"
	"net/http"
	"qkeruen/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWTDriver(jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "No toke found."})
			return
		}

		token, err := jwtService.ValidateToken(authHeader)

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			if claims["role"] != "driver" {
				c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": "wrong type role"})
				return
			}
			log.Println(claims["phone_number"])
			log.Println(claims["role"])
		} else {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": "unauthorized"})
		}

	}
}

func AuthorizeJWTUser(jwtService service.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "No toke found."})
			return
		}

		token, err := jwtService.ValidateToken(authHeader)

		if token.Valid {
			claims := token.Claims.(jwt.MapClaims)
			if claims["role"] != "user" {
				c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": "wrong type role"})
				return
			}
			log.Println(claims["phone_number"])
			log.Println(claims["role"])
		} else {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": "unauthorized"})
		}

	}
}
