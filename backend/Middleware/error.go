package middleware

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		println("Testststst")
		if len(c.Errors) > 0 {
			err := c.Errors.Last()
			log.Println("Error occurred:", err.Error())
			statusCode, isExisting := c.Get("statusCode")

			if isExisting {
				c.JSON(statusCode.(int), gin.H{
					"error": err.Error(),
				})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Something went wrong",
				})
			}
			return
		}
	}
}
