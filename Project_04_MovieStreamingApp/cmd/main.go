package main

import (
	"fmt"
	"log"
	"movie-streaming-app/internal/config"
	"movie-streaming-app/internal/database"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config", err)
	}

	pool, err := database.ConnectDb(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database", err)
	}

	defer pool.Close()

	var router = gin.Default()
	err = router.SetTrustedProxies(nil)
	if err != nil {
		fmt.Println("SetTrustedProxies err: ", err)
		return
	}

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.GET("/check-health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "It's works!",
		})
	})

	err = router.Run(":" + cfg.Port)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		return
	}
}
