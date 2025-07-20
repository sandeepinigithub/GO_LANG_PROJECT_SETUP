package main

import (
	"log"
	"net/http"
	"time"
	"devsMailGo/config"
	"devsMailGo/logger"
	"devsMailGo/middleware"
	"devsMailGo/routes"
)

func main() {
	// Initialize logger
	logger.InitLogger("devsMailGo-API", "info", "json")
	logger.GlobalLogger.Info("Starting devsMailGo API")

	// Load configuration
	if err := config.LoadConfig(); err != nil {
		logger.GlobalLogger.Error("Failed to load configuration", err)
		log.Fatal("Configuration error:", err)
	}

	// Initialize database
	if err := config.InitDB(); err != nil {
		logger.GlobalLogger.Error("Failed to connect to database", err)
		log.Fatal("Database connection error:", err)
	}
	logger.GlobalLogger.Info("Database connected successfully")

	// Setup routes
	router := routes.SetupRoutes()

	// Apply middleware stack
	handler := middleware.SecureHeadersMiddleware(router)
	handler = middleware.CORSMiddleware(handler)
	
	// Apply rate limiting if configured
	if config.AppConfig != nil && config.AppConfig.Security.RateLimitRequests > 0 {
		handler = middleware.RateLimitMiddleware(
			config.AppConfig.Security.RateLimitRequests,
			config.AppConfig.Security.RateLimitWindow,
		)(handler)
	}

	// Configure server
	server := &http.Server{
		Addr:         ":" + config.AppConfig.Server.Port,
		Handler:      handler,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	// Start server
	logger.GlobalLogger.Info("Server starting", map[string]interface{}{
		"port":        config.AppConfig.Server.Port,
		"environment": config.AppConfig.Server.Environment,
	})
	
	log.Printf("Server starting on port %s", config.AppConfig.Server.Port)
	log.Printf("Health check available at: http://localhost:%s/api/health", config.AppConfig.Server.Port)
	log.Printf("Login endpoint: http://localhost:%s/api/login", config.AppConfig.Server.Port)
	
	if err := server.ListenAndServe(); err != nil {
		logger.GlobalLogger.Error("Failed to start server", err)
		log.Fatal("Server error:", err)
	}
}
