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
			"service": "gateway",
		})
	})

	// Proxy routes - all requests go through gateway
	api := r.Group("/api/v1")
	{
		// Auth routes -> auth-service
		api.Any("/auth/*path", proxyToService("AUTH_SERVICE_URL"))

		// Account routes -> account-service
		api.Any("/users/*path", proxyToService("ACCOUNT_SERVICE_URL"))
		api.Any("/orgs/*path", proxyToService("ACCOUNT_SERVICE_URL"))

		// Agent routes -> agent-service
		api.Any("/agents/*path", proxyToService("AGENT_SERVICE_URL"))
		api.Any("/voices/*path", proxyToService("AGENT_SERVICE_URL"))

		// Billing routes -> billing-service
		api.Any("/billing/*path", proxyToService("BILLING_SERVICE_URL"))

		// Numbers routes -> numbers-service
		api.Any("/numbers/*path", proxyToService("NUMBERS_SERVICE_URL"))

		// Webhooks routes -> webhooks-service
		api.Any("/webhooks/*path", proxyToService("WEBHOOKS_SERVICE_URL"))

		// Calls/CDR routes -> cdr-service
		api.Any("/calls/*path", proxyToService("CDR_SERVICE_URL"))

		// Rooms routes -> orchestrator
		api.Any("/rooms/*path", proxyToService("ORCHESTRATOR_URL"))

		// Analytics routes -> analytics-service
		api.Any("/analytics/*path", proxyToService("ANALYTICS_SERVICE_URL"))
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	r.Run(":" + port)
}

// proxyToService returns a handler that proxies requests to the given service
func proxyToService(envVar string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// TODO: Implement actual proxy logic
		// 1. Validate JWT token
		// 2. Check rate limits
		// 3. Forward request to downstream service
		c.JSON(http.StatusNotImplemented, gin.H{
			"error":   "proxy not implemented",
			"service": envVar,
			"path":    c.Request.URL.Path,
		})
	}
}
