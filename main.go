package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := setupRouter()
	_ = router.Run(":8080")
}

func setupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Hello World!")
	})

	return router
}
