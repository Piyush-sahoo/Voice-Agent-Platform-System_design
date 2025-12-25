import express, { Request, Response } from 'express';
import cors from 'cors';
import helmet from 'helmet';

const app = express();
const PORT = process.env.PORT || 3003;

// Middleware
app.use(helmet());
app.use(cors());
app.use(express.json());

// Health check
app.get('/health', (req: Request, res: Response) => {
    res.json({ status: 'healthy', service: 'numbers-service' });
});

// ===================
// SEARCH & BUY ENDPOINTS
// ===================

// Search available numbers
app.get('/api/v1/numbers/available', (req: Request, res: Response) => {
    // TODO: Search Twilio/Telnyx for available numbers
    res.status(501).json({ endpoint: 'GET /api/v1/numbers/available' });
});

// Purchase number
app.post('/api/v1/numbers/buy', (req: Request, res: Response) => {
    // TODO: Purchase number from provider
    res.status(501).json({ endpoint: 'POST /api/v1/numbers/buy' });
});

// ===================
// INVENTORY ENDPOINTS
// ===================

// List org numbers
app.get('/api/v1/numbers', (req: Request, res: Response) => {
    // TODO: List organization's phone numbers
    res.status(501).json({ endpoint: 'GET /api/v1/numbers' });
});

// Get number details
app.get('/api/v1/numbers/:id', (req: Request, res: Response) => {
    // TODO: Get number details
    res.status(501).json({ endpoint: 'GET /api/v1/numbers/:id' });
});

// Release number
app.delete('/api/v1/numbers/:id', (req: Request, res: Response) => {
    // TODO: Release number back to provider
    res.status(501).json({ endpoint: 'DELETE /api/v1/numbers/:id' });
});

// ===================
// ASSIGNMENT ENDPOINTS
// ===================

// Assign to agent
app.put('/api/v1/numbers/:id/assign', (req: Request, res: Response) => {
    // TODO: Assign number to agent
    res.status(501).json({ endpoint: 'PUT /api/v1/numbers/:id/assign' });
});

// Unassign
app.delete('/api/v1/numbers/:id/assign', (req: Request, res: Response) => {
    // TODO: Unassign number from agent
    res.status(501).json({ endpoint: 'DELETE /api/v1/numbers/:id/assign' });
});

// ===================
// ROUTING ENDPOINTS
// ===================

// Set routing rules
app.put('/api/v1/numbers/:id/routing', (req: Request, res: Response) => {
    // TODO: Set call routing rules
    res.status(501).json({ endpoint: 'PUT /api/v1/numbers/:id/routing' });
});

// ===================
// PORTING ENDPOINTS
// ===================

// Port number request
app.post('/api/v1/numbers/port', (req: Request, res: Response) => {
    // TODO: Submit port request
    res.status(501).json({ endpoint: 'POST /api/v1/numbers/port' });
});

// Port status
app.get('/api/v1/numbers/port/:id', (req: Request, res: Response) => {
    // TODO: Get port request status
    res.status(501).json({ endpoint: 'GET /api/v1/numbers/port/:id' });
});

// Start server
app.listen(PORT, () => {
    console.log(`Numbers service running on port ${PORT}`);
});
