package main

import (
	"log"
	"os"

	middleware "golang-restaurant-management/middleware"
	routes "golang-restaurant-management/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// Load environment variables from .env file if present
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables from the system")
	} else {
		log.Println(".env file loaded successfully")
	}

}

func main() {
	// Read the PORT environment variable
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port for local development
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery()) // Add Recovery middleware to handle panics

	// Set up routes
	routes.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)

	// Start the server
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
