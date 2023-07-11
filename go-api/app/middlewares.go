package app

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

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
