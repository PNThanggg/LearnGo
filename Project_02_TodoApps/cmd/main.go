package main

import (
	"fmt"
	"log"
	"todo-apps/internal/config"
	"todo-apps/internal/database"
	"todo-apps/internal/handlers"

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
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status":   "success",
			"message":  "Todo API is running",
			"database": "Connected",
		})
	})

	router.GET("/todos", handlers.GetAllTodosHandler(pool))
	router.POST("/todos", handlers.CreateTodoHandler(pool))
	router.GET("/todos/:id", handlers.GetTodoHandler(pool))

	err = router.Run(":" + cfg.Port)
	if err != nil {
		fmt.Printf("Error starting server: %s\n", err)
		return
	}
}
