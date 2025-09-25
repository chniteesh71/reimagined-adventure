package main

import (
	"fancy-adventure/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Serve static files
	r.Static("/static", "./static")
	// Load templates
	r.LoadHTMLGlob("templates/*")

	// Routes
	r.GET("/", handlers.Home)
	r.GET("/adventure/:name", handlers.Adventure)

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
