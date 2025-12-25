package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Health check
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "healthy",
			"service": "sip-ingress",
		})
	})

	// SIP trunk management
	trunks := r.Group("/api/v1/trunks")
	{
		// List SIP trunks
		trunks.GET("", handleListTrunks)

		// Create SIP trunk
		trunks.POST("", handleCreateTrunk)

		// Get trunk details
		trunks.GET("/:id", handleGetTrunk)

		// Update trunk
		trunks.PUT("/:id", handleUpdateTrunk)

		// Delete trunk
		trunks.DELETE("/:id", handleDeleteTrunk)
	}

	// Config management
	config := r.Group("/api/v1/config")
	{
		// Reload SIP config
		config.POST("/reload", handleReloadConfig)

		// Get current config
		config.GET("", handleGetConfig)
	}

	// Dispatch rules
	dispatch := r.Group("/api/v1/dispatch-rules")
	{
		// List dispatch rules
		dispatch.GET("", handleListDispatchRules)

		// Create dispatch rule
		dispatch.POST("", handleCreateDispatchRule)

		// Delete dispatch rule
		dispatch.DELETE("/:id", handleDeleteDispatchRule)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8007"
	}
	r.Run(":" + port)
}

// ===================
// HANDLERS (STUBS)
// ===================

func handleListTrunks(c *gin.Context) {
	// TODO: List SIP trunks from LiveKit
	c.JSON(http.StatusNotImplemented, gin.H{"endpoint": "GET /api/v1/trunks"})
}

func handleCreateTrunk(c *gin.Context) {
	// TODO: Create SIP trunk in LiveKit
	c.JSON(http.StatusNotImplemented, gin.H{"endpoint": "POST /api/v1/trunks"})
}

func handleGetTrunk(c *gin.Context) {
	// TODO: Get trunk details
	c.JSON(http.StatusNotImplemented, gin.H{"endpoint": "GET /api/v1/trunks/:id"})
}

func handleUpdateTrunk(c *gin.Context) {
	// TODO: Update SIP trunk
	c.JSON(http.StatusNotImplemented, gin.H{"endpoint": "PUT /api/v1/trunks/:id"})
}

func handleDeleteTrunk(c *gin.Context) {
	// TODO: Delete SIP trunk
	c.JSON(http.StatusNotImplemented, gin.H{"endpoint": "DELETE /api/v1/trunks/:id"})
}

func handleReloadConfig(c *gin.Context) {
	// TODO: Reload SIP configuration
	c.JSON(http.StatusNotImplemented, gin.H{"endpoint": "POST /api/v1/config/reload"})
}

func handleGetConfig(c *gin.Context) {
	// TODO: Get current SIP config
	c.JSON(http.StatusNotImplemented, gin.H{"endpoint": "GET /api/v1/config"})
}

func handleListDispatchRules(c *gin.Context) {
	// TODO: List dispatch rules
	c.JSON(http.StatusNotImplemented, gin.H{"endpoint": "GET /api/v1/dispatch-rules"})
}

func handleCreateDispatchRule(c *gin.Context) {
	// TODO: Create dispatch rule
	c.JSON(http.StatusNotImplemented, gin.H{"endpoint": "POST /api/v1/dispatch-rules"})
}

func handleDeleteDispatchRule(c *gin.Context) {
	// TODO: Delete dispatch rule
	c.JSON(http.StatusNotImplemented, gin.H{"endpoint": "DELETE /api/v1/dispatch-rules/:id"})
}
