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
			"service": "billing-service",
		})
	})

	// Public billing endpoints
	billing := r.Group("/api/v1/billing")
	{
		// Balance
		billing.GET("/balance", handleGetBalance)

		// Top-up / Add credits
		billing.POST("/top-up", handleTopUp)

		// Usage
		billing.GET("/usage", handleGetUsage)
		billing.GET("/usage/daily", handleGetDailyUsage)

		// Invoices
		billing.GET("/invoices", handleListInvoices)
		billing.GET("/invoices/:id/pdf", handleDownloadInvoice)

		// Subscription
		billing.GET("/subscription", handleGetSubscription)
		billing.POST("/subscription", handleUpdateSubscription)
		billing.DELETE("/subscription", handleCancelSubscription)
	}

	// Internal billing endpoints (for orchestrator/agent-worker)
	internal := r.Group("/internal/v1")
	{
		// Billing lease (heartbeat system)
		internal.POST("/lease", handleRequestLease)
		internal.POST("/lease/renew", handleRenewLease)
		internal.POST("/lease/release", handleReleaseLease)

		// Deduct credits
		internal.POST("/deduct", handleDeductCredits)
	}

	// Stripe webhooks
	r.POST("/webhooks/stripe", handleStripeWebhook)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8002"
	}
	r.Run(":" + port)
}

// ===================
// HANDLERS (STUBS)
// ===================

func handleGetBalance(c *gin.Context) {
	// TODO: Get org balance from database
	c.JSON(http.StatusNotImplemented, gin.H{"endpoint": "GET /api/v1/billing/balance"})
}

func handleTopUp(c *gin.Context) {
	// TODO: Add credits via Stripe payment
	c.JSON(http.StatusNotImplemented, gin.H{"endpoint": "POST /api/v1/billing/top-up"})
}

func handleGetUsage(c *gin.Context) {
	// TODO: Get usage summary
	c.JSON(http.StatusNotImplemented, gin.H{"endpoint": "GET /api/v1/billing/usage"})
}

func handleGetDailyUsage(c *gin.Context) {
	// TODO: Get daily usage breakdown
	c.JSON(http.StatusNotImplemented, gin.H{"endpoint": "GET /api/v1/billing/usage/daily"})
}

func handleListInvoices(c *gin.Context) {
	// TODO: List invoices
	c.JSON(http.StatusNotImplemented, gin.H{"endpoint": "GET /api/v1/billing/invoices"})
}

func handleDownloadInvoice(c *gin.Context) {
	// TODO: Generate and download invoice PDF
	c.JSON(http.StatusNotImplemented, gin.H{"endpoint": "GET /api/v1/billing/invoices/:id/pdf"})
}

func handleGetSubscription(c *gin.Context) {
	// TODO: Get current subscription plan
	c.JSON(http.StatusNotImplemented, gin.H{"endpoint": "GET /api/v1/billing/subscription"})
}

func handleUpdateSubscription(c *gin.Context) {
	// TODO: Change subscription plan
	c.JSON(http.StatusNotImplemented, gin.H{"endpoint": "POST /api/v1/billing/subscription"})
}

func handleCancelSubscription(c *gin.Context) {
	// TODO: Cancel subscription
	c.JSON(http.StatusNotImplemented, gin.H{"endpoint": "DELETE /api/v1/billing/subscription"})
}

func handleRequestLease(c *gin.Context) {
	// TODO: Request billing lease for call (reserve credits)
	c.JSON(http.StatusNotImplemented, gin.H{"endpoint": "POST /internal/v1/lease"})
}

func handleRenewLease(c *gin.Context) {
	// TODO: Renew billing lease (heartbeat)
	c.JSON(http.StatusNotImplemented, gin.H{"endpoint": "POST /internal/v1/lease/renew"})
}

func handleReleaseLease(c *gin.Context) {
	// TODO: Release billing lease (call ended)
	c.JSON(http.StatusNotImplemented, gin.H{"endpoint": "POST /internal/v1/lease/release"})
}

func handleDeductCredits(c *gin.Context) {
	// TODO: Deduct credits for completed call
	c.JSON(http.StatusNotImplemented, gin.H{"endpoint": "POST /internal/v1/deduct"})
}

func handleStripeWebhook(c *gin.Context) {
	// TODO: Handle Stripe webhook events
	c.JSON(http.StatusNotImplemented, gin.H{"endpoint": "POST /webhooks/stripe"})
}
