package main

import (
	"fmt"
	"go_tweets/internal/config"
	"go_tweets/internal/database"
	"go_tweets/internal/handler"
	"go_tweets/internal/repository"
	"go_tweets/internal/service"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Failed to load config", err)
		return
	}

	pool, err := database.ConnectDb(cfg.DatabaseURL)
	if err != nil {
		log.Fatal("Failed to connect to database", err)
		return
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

	userRepository := repository.NewUserRepository(pool)
	userService := service.NewUserService(cfg, userRepository)
	userHandler := handler.NewHandler(router, userService)
	userHandler.RouterList()

	err = router.Run(":" + cfg.Port)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		return
	}
}
