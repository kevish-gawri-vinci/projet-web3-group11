package middleware

import (
	utils "backend/Utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

//For protected routes ( routes that need user to be authentified )

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		println("In the auth middleware")
		tokenString := c.GetHeader("Authorization")

		// Parse the token
		token, err := utils.VerifyToken(tokenString)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			c.Abort()
			return
		}

		// Set the token claims to the context
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("claims", claims)
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		println("Token validated, calling c.Next()")
		c.Next() // Proceed to the next handler if authorized
	}
}
