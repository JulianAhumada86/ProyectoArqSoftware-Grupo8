package app

import (
	"net/http"
	"time"

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
				// Obtener la fecha de creación del token como Unix timestamp
				creationTime, ok := claims["fecha"].(float64)
				if !ok {
					c.JSON(http.StatusUnauthorized, gin.H{"error": "No se pudo obtener la fecha de creación del token"})
					c.Abort()
					return
				}

				// Convertir el Unix timestamp a time.Time
				creationDate := time.Unix(int64(creationTime), 0)

				// Verificar si ha pasado más de un día
				if time.Since(creationDate).Hours() <= 6 {
					isAdmin, ok := claims["admin"].(float64)
					if ok && isAdmin == 1 {
						c.Next()
					} else {
						c.JSON(http.StatusUnauthorized, gin.H{"error": "Tienes que ser administrador para esta tarea"})
						c.Abort()
					}
				} else {
					c.JSON(http.StatusUnauthorized, gin.H{"error": "El token ha expirado"})
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
			claims, ok := token.Claims.(jwt.MapClaims)
			if ok {
				// Obtener la fecha de creación del token como Unix timestamp
				creationTime, ok := claims["fecha"].(float64)
				if !ok {
					c.JSON(http.StatusUnauthorized, gin.H{"error": "No se pudo obtener la fecha de creación del token"})
					c.Abort()
					return
				}

				// Convertir el Unix timestamp a time.Time
				creationDate := time.Unix(int64(creationTime), 0)

				// Verificar si ha pasado más de un día
				if time.Since(creationDate).Hours() <= 24 {
					c.Next()
					return
				}
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token expiro"})
				c.Abort()
				return
			}
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Token de autorización inválido"})
			c.Abort()
			return
		}
	}
}
