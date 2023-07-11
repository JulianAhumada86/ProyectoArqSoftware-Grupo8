package app

import (
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AdminTokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token de autorización faltante"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("Secret key"), nil // Clave secreta para verificar el token
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token de autorización inválido"})
			c.Abort()
			return
		}

		if token.Valid {
			claims, ok := token.Claims.(jwt.MapClaims)
			if ok {
				isAdmin, ok := claims["admin"].(float64)
				log.Println(isAdmin)
				log.Println("esadmin")

				if ok && isAdmin == 1 {
					c.Next()
				} else {
					c.JSON(http.StatusUnauthorized, gin.H{"error": "Tenes que ser administrador para esta tarea"})
					c.Abort()
				}
			}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token de autorización inválido"})
			c.Abort()
			return
		}
	}
}

func TokenMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token de autorización faltante"})
			c.Abort()
			return
		}

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte("Secret key"), nil // Clave secreta para verificar el token
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token de autorización inválido"})
			c.Abort()
			return
		}

		if token.Valid {
			// Token válido, continuar con la solicitud
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token de autorización inválido"})
			c.Abort()
			return
		}
	}
}
