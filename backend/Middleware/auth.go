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
		if len(tokenString) == 0 {
			utils.ThrowError(c, &utils.ErrorStruct{Msg: "No token in headers", Code: http.StatusUnauthorized})
			c.Abort()
			return
		}

		// Parse the token
		token, err := utils.VerifyToken(tokenString)

		if err != nil {
			utils.ThrowError(c, &utils.ErrorStruct{Msg: err.Error(), Code: http.StatusUnauthorized})
			c.Abort()
			return
		}

		// Set the token claims to the context
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			c.Set("claims", claims)
		} else {
			utils.ThrowError(c, &utils.ErrorStruct{Msg: "Unauthorized", Code: http.StatusUnauthorized})
			c.Abort()
			return
		}
		println("Token validated, calling c.Next()")
		c.Next() // Proceed to the next handler if authorized
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		println("Check if user is admin ")
		isAdmin := utils.GetUserRoleInClaims(ctx)
		if isAdmin {
			ctx.Next()
		} else {
			utils.ThrowError(ctx, &utils.ErrorStruct{Msg: "User is not admin", Code: http.StatusUnauthorized})
			ctx.Abort()
		}
	}
}
