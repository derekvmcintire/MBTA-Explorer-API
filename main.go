package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/api", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello, MBTA Train Tracker API!",
		})
	})

	r.Run(":8080") // Listen on port 8080
}
