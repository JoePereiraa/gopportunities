package router

import (
	"github.com/gin-gonic/gin"
)

func Initialize() {
	// Create a Gin router with default middleware (logger and recovery)
	router := gin.Default()

	// Initializes routes
	InitializeRoutes(router)

	// Start server on port 8080 (default)
	// Server will listen on 0.0.0.0:8080 (localhost:8080 on Windows)
	router.Run()
}
