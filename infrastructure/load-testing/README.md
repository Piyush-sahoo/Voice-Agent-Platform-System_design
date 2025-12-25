# Load Testing

Scripts for load testing the VAP Platform.

## Tools

- **K6** - HTTP endpoint load testing
- **Locust** - Python-based load testing

## Running Tests

```bash
# K6 - Test API Gateway
k6 run scripts/gateway-load.js

# K6 - Simulate 1000 concurrent calls
k6 run --vus 1000 --duration 5m scripts/call-simulation.js
```
