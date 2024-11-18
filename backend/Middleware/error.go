package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			log.Println("Error occurred:", err.Error())
			statusCode, isExisting := c.Get("statusCode")

			if isExisting {
				c.JSON(statusCode.(int), gin.H{
					"error": err.Error(),
				})
				c.Abort()
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Something went wrong",
				})
				c.Abort()
			}
			return
		}
	}
}
