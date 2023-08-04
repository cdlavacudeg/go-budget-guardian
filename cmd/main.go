package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/cdlavacudeg/go-budget-guardian/config"
)

func main() {
	_, err := config.LoadConfig()
	if err != nil {
		fmt.Println("Error loading the configuration", err)
		return
	}

	router := setupRouter()
	_ = router.Run(":8080")
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	router.GET("ping", func(c *gin.Context) {
		cfg, err := config.GetConfig()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, cfg)
	})
	return router
}
