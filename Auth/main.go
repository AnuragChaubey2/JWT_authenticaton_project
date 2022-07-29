package main

import (
	"github.com/Auth/routes"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	port := os.Getenv("PORT")
	if port == " " {
		port = "8000"
	}
	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)
	routes.UserRoutes(router)

	// api end points
	//no ned to pass response and req all we have to passs is gin and it will handle
	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "Acess granted for api-1"})
	})
	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(200, gin.H{"success": "access granted for api-2"})
	})

	router.Run(":" + port)
}
