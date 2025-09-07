# Fundament Application Tasks
.PHONY: start stop restart status logs clean reset help db-logs backend-logs frontend-logs

# Default target
help:
	@echo "🚀 Fundament Application Commands:"
	@echo ""
	@echo "📦 Application Management:"
	@echo "  make start     - Start all services (builds and runs docker-compose)"
	@echo "  make stop      - Stop all services"
	@echo "  make restart   - Restart all services"
	@echo "  make reset     - Reset everything (stop, remove volumes, start fresh)"
	@echo ""
	@echo "🔍 Monitoring:"
	@echo "  make status    - Show status of all services"
	@echo "  make logs      - Show logs from all services"
	@echo "  make db-logs   - Show database logs only"
	@echo "  make backend-logs - Show backend logs only"
	@echo "  make frontend-logs - Show frontend logs only"
	@echo ""
	@echo "🧹 Cleanup:"
	@echo "  make clean     - Remove all containers and networks (keeps data)"
	@echo ""
	@echo "📋 URLs when running:"
	@echo "  Frontend: http://localhost:3000"
	@echo "  Backend API: http://localhost:8080"

# Launch the application
start:
	@echo "🚀 Starting Fundament application..."
	@echo "This may take a minute to build and start all services..."
	docker-compose up --build -d
	@echo ""
	@echo "⏳ Waiting for services to be healthy..."
	@timeout /t 5 >nul 2>&1 || ping -n 6 127.0.0.1 >nul
	@echo ""
	@echo "✅ Application launched successfully!"
	@echo "🌐 Frontend: http://localhost:3000"
	@echo "🔧 Backend API: http://localhost:8080"
	@echo "💾 Database: localhost:5432 (internal)"
	@echo ""
	@echo "💡 Tip: Use 'make status' to check if all services are healthy"

# Stop the application
stop:
	@echo "🛑 Stopping Fundament application..."
	docker-compose down
	@echo "✅ Application stopped successfully"

# Restart all services
restart:
	@echo "🔄 Restarting Fundament application..."
	docker-compose restart
	@echo "✅ Application restarted successfully"

# Reset everything (nuclear option)
reset:
	@echo "💥 Resetting Fundament application (this will delete all data!)..."
	@echo "⚠️  This will remove ALL data including users and notes!"
	@echo "Press Ctrl+C to cancel, or press Enter to continue..."
	@timeout /t 10 >nul 2>&1 || ping -n 11 127.0.0.1 >nul
	docker-compose down -v
	docker system prune -f
	@echo "✅ Everything reset! Run 'make start' to begin fresh."

# Show status of all services
status:
	@echo "📊 Fundament Application Status:"
	@echo "================================="
	docker-compose ps

# Show logs from all services
logs:
	@echo "📋 Recent logs from all services:"
	@echo "=================================="
	docker-compose logs --tail=20

# Show database logs only
db-logs:
	@echo "💾 Database logs:"
	@echo "================="
	docker-compose logs postgres --tail=20

# Show backend logs only
backend-logs:
	@echo "🔧 Backend logs:"
	@echo "================"
	docker-compose logs backend --tail=20

# Show frontend logs only
frontend-logs:
	@echo "🌐 Frontend logs:"
	@echo "================="
	docker-compose logs frontend --tail=20

# Clean up containers and networks (keeps data volumes)
clean:
	@echo "🧹 Cleaning up containers and networks..."
	docker-compose down
	docker system prune -f
	@echo "✅ Cleanup completed"
