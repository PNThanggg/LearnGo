package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	var router = gin.Default()
	router.SetTrustedProxies(nil)
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "Todo API is running",
		})
	})

	router.Run(":8080")
}
