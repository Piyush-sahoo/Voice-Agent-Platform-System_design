.PHONY: build up down logs health

# Build all services
build:
	docker-compose build

# Start all services
up:
	docker-compose up -d

# Stop all services
down:
	docker-compose down

# View logs
logs:
	docker-compose logs -f

# Health check all services
health:
	@echo "Checking Gateway..."
	@curl -s http://localhost:8000/health || echo "Gateway: DOWN"
	@echo "Checking Auth..."
	@curl -s http://localhost:8001/health || echo "Auth: DOWN"
	@echo "Checking Billing..."
	@curl -s http://localhost:8002/health || echo "Billing: DOWN"
	@echo "Checking CDR..."
	@curl -s http://localhost:8003/health || echo "CDR: DOWN"
	@echo "Checking Orchestrator..."
	@curl -s http://localhost:8004/health || echo "Orchestrator: DOWN"
	@echo "Checking Analytics..."
	@curl -s http://localhost:8005/health || echo "Analytics: DOWN"
	@echo "Checking Account..."
	@curl -s http://localhost:3001/health || echo "Account: DOWN"
	@echo "Checking Agent..."
	@curl -s http://localhost:3002/health || echo "Agent: DOWN"
	@echo "Checking Numbers..."
	@curl -s http://localhost:3003/health || echo "Numbers: DOWN"
	@echo "Checking Webhooks..."
	@curl -s http://localhost:3004/health || echo "Webhooks: DOWN"

# Generate protobuf
proto:
	cd libs/protos && buf generate
