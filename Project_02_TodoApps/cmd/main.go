package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	var router = gin.Default()
	err := router.SetTrustedProxies(nil)
	if err != nil {
		fmt.Println("SetTrustedProxies err: ", err)
		return
	}
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":  "success",
			"message": "Todo API is running",
		})
	})

	err = router.Run(":8080")
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		return
	}
}
