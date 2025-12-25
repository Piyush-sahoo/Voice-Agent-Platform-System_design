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
			"service": "orchestrator",
		})
	})

	// Room management endpoints
	rooms := r.Group("/api/v1/rooms")
	{
		// Create room
		rooms.POST("", handleCreateRoom)

		// Get room status
		rooms.GET("/:id", handleGetRoom)

		// Close room
		rooms.DELETE("/:id", handleCloseRoom)
	}

	// Call management endpoints
	calls := r.Group("/api/v1/calls")
	{
		// Initiate outbound call
		calls.POST("/outbound", handleOutboundCall)

		// Transfer call
		calls.POST("/:id/transfer", handleTransferCall)

		// End call
		calls.POST("/:id/end", handleEndCall)

		// Get join token (WebRTC)
		calls.GET("/:id/token", handleGetJoinToken)
	}

	// Internal endpoints
	internal := r.Group("/internal/v1")
	{
		// Dispatch agent worker
		internal.POST("/dispatch", handleDispatch)

		// List active workers
		internal.GET("/workers", handleListWorkers)

		// Worker heartbeat
		internal.POST("/workers/:id/heartbeat", handleWorkerHeartbeat)

		// Room watchdog (check for zombie rooms)
		internal.POST("/watchdog", handleWatchdog)
	}

	// WebSocket endpoint for real-time events
	r.GET("/ws", handleWebSocket)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8004"
	}
	r.Run(":" + port)
}

// ===================
// HANDLERS (STUBS)
// ===================

func handleCreateRoom(c *gin.Context) {
	// TODO: Create LiveKit room
	c.JSON(http.StatusNotImplemented, gin.H{"endpoint": "POST /api/v1/rooms"})
}

func handleGetRoom(c *gin.Context) {
	// TODO: Get room status
	c.JSON(http.StatusNotImplemented, gin.H{"endpoint": "GET /api/v1/rooms/:id"})
}

func handleCloseRoom(c *gin.Context) {
	// TODO: Close room and disconnect participants
	c.JSON(http.StatusNotImplemented, gin.H{"endpoint": "DELETE /api/v1/rooms/:id"})
}

func handleOutboundCall(c *gin.Context) {
	// TODO: Initiate outbound call via SIP
	c.JSON(http.StatusNotImplemented, gin.H{"endpoint": "POST /api/v1/calls/outbound"})
}

func handleTransferCall(c *gin.Context) {
	// TODO: Transfer call to another number/agent
	c.JSON(http.StatusNotImplemented, gin.H{"endpoint": "POST /api/v1/calls/:id/transfer"})
}

func handleEndCall(c *gin.Context) {
	// TODO: End active call
	c.JSON(http.StatusNotImplemented, gin.H{"endpoint": "POST /api/v1/calls/:id/end"})
}

func handleGetJoinToken(c *gin.Context) {
	// TODO: Generate LiveKit join token
	c.JSON(http.StatusNotImplemented, gin.H{"endpoint": "GET /api/v1/calls/:id/token"})
}

func handleDispatch(c *gin.Context) {
	// TODO: Dispatch agent worker to room
	c.JSON(http.StatusNotImplemented, gin.H{"endpoint": "POST /internal/v1/dispatch"})
}

func handleListWorkers(c *gin.Context) {
	// TODO: List active agent workers
	c.JSON(http.StatusNotImplemented, gin.H{"endpoint": "GET /internal/v1/workers"})
}

func handleWorkerHeartbeat(c *gin.Context) {
	// TODO: Handle worker heartbeat
	c.JSON(http.StatusNotImplemented, gin.H{"endpoint": "POST /internal/v1/workers/:id/heartbeat"})
}

func handleWatchdog(c *gin.Context) {
	// TODO: Check for zombie rooms (no participants)
	c.JSON(http.StatusNotImplemented, gin.H{"endpoint": "POST /internal/v1/watchdog"})
}

func handleWebSocket(c *gin.Context) {
	// TODO: WebSocket handler for real-time events
	c.JSON(http.StatusNotImplemented, gin.H{"endpoint": "GET /ws"})
}
