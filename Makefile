# Fundament Application Tasks
.PHONY: start stop help

# Default target
help:
	@echo "Available commands:"
	@echo "  make start  - Start the application (builds and runs docker-compose services)"
	@echo "  make stop    - Stop the application and clean up containers"

# Launch the application
start:
	@echo " Launching Fundament application..."
	docker-compose up --build -d
	@echo " Application launched!"
	@echo " Frontend: http://localhost:3000"
	@echo " Backend API: http://localhost:8080"
	@echo "Database: localhost:5432"

# Kill the application
stop:
	@echo "🛑 Stopping Fundament application..."
	docker-compose down -v
	@echo "✅ Application stopped and containers cleaned up"
