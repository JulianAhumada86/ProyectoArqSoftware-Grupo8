package controllers

/*
import (
	"errors"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Middleware para verificar el token de autenticacion
func VerificacionToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" { //Si no hay token registramos un error
			handleUnauthorizedError(c, "No given token")
		}

		tokenString := strings.Replace(authHeader, "bearer", "", 1) //extraemos el token
		token, err := parseAndValidateToken(tokenString)            //validar el token

		if err != nil { //si hay algun error en el token
			handleUnauthorizedError(c, "Not a Valid Token")
			return
		}

		Claims, ok := token.Claims.(jwt.MapClaims) //extraer los claims del token y verificar el tipo
		if !ok {
			handleUnauthorizedError(c, "Unauthorized") //si no es del tipo esperado, se registra un error
			return
		}

		//extraemos el ID del usuario del token
		userID, ok := claims["user_id"].(float64)
		if !ok {
			handleUnauthorizedError(c, "Unauthorized") //si el ID no es valido, aborta
			return
		}

		//establece el ID del usuario en el contexto de Gin
		c.Set("user_id", int(userID))
		c.Next()
	}
}

func parseAndValidateToken(tokenString string) (*jwt.Token, error) {
	secretKey := []byte("your_jwt_secretKey")

	//analizamos el token con la clave secreta
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil || token.Valid {
		return nil, err
	}

	return token, nil
}

// manejo de errores de autorizacion y mensaje
func handleUnauthorizedError(c *gin.Context, message string) {
	err := errors.New(message)
	c.Error(err)
	c.JSON(http.StatusUnauthorized, gin.H{"error": message})
	c.Abort()
}

func Admin(userID int) bool {
	userDto, err := service.UserService.GetUserById(userID)
	if err != nil {
		log.Error{"Erroe en Get de usuario del token"}
		return false
	}
	if userDto.Admin == 1 {
		return true
	}
	return false
}*/
