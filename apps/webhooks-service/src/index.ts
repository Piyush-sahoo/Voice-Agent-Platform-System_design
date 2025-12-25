import express, { Request, Response } from 'express';
import cors from 'cors';
import helmet from 'helmet';

const app = express();
const PORT = process.env.PORT || 3004;

// Middleware
app.use(helmet());
app.use(cors());
app.use(express.json());

// Health check
app.get('/health', (req: Request, res: Response) => {
    res.json({ status: 'healthy', service: 'webhooks-service' });
});

// ===================
// WEBHOOK CONFIG ENDPOINTS
// ===================

// Create webhook endpoint
app.post('/api/v1/webhooks', (req: Request, res: Response) => {
    // TODO: Create new webhook endpoint
    res.status(501).json({ endpoint: 'POST /api/v1/webhooks' });
});

// List webhooks
app.get('/api/v1/webhooks', (req: Request, res: Response) => {
    // TODO: List webhook endpoints
    res.status(501).json({ endpoint: 'GET /api/v1/webhooks' });
});

// Update webhook
app.put('/api/v1/webhooks/:id', (req: Request, res: Response) => {
    // TODO: Update webhook endpoint
    res.status(501).json({ endpoint: 'PUT /api/v1/webhooks/:id' });
});

// Delete webhook
app.delete('/api/v1/webhooks/:id', (req: Request, res: Response) => {
    // TODO: Delete webhook endpoint
    res.status(501).json({ endpoint: 'DELETE /api/v1/webhooks/:id' });
});

// ===================
// EVENT ENDPOINTS
// ===================

// Event history
app.get('/api/v1/webhooks/:id/events', (req: Request, res: Response) => {
    // TODO: Get webhook event history
    res.status(501).json({ endpoint: 'GET /api/v1/webhooks/:id/events' });
});

// Send test event
app.post('/api/v1/webhooks/:id/test', (req: Request, res: Response) => {
    // TODO: Send test webhook event
    res.status(501).json({ endpoint: 'POST /api/v1/webhooks/:id/test' });
});

// ===================
// INTERNAL ENDPOINTS
// ===================

// Dispatch webhook (from Kafka)
app.post('/internal/v1/dispatch', (req: Request, res: Response) => {
    // TODO: Dispatch webhook with retry logic
    res.status(501).json({ endpoint: 'POST /internal/v1/dispatch' });
});

// Retry failed webhook
app.post('/internal/v1/retry/:id', (req: Request, res: Response) => {
    // TODO: Retry failed webhook
    res.status(501).json({ endpoint: 'POST /internal/v1/retry/:id' });
});

// Start server
app.listen(PORT, () => {
    console.log(`Webhooks service running on port ${PORT}`);
});
