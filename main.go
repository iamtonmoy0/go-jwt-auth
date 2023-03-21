package main

import (
	"os"

	"github.com/iamtonmoy0/go-jwt-auth/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	port = os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())
	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"Success": "Access Granted For api-1"})
	})
	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"Success": "Access Granted For api-2"})
	})

	router.Run(":" + port)

}
