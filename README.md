# 📝 Fundament

A modern, full-stack note-taking application built with Go, React, and PostgreSQL. Features user authentication, real-time note management, and a clean, responsive interface.

## 🚀 Features

- **User Authentication**: Secure JWT-based authentication system
- **Note Management**: Create, read, update, and delete notes
- **Responsive Design**: Clean, modern UI that works on all devices
- **Real-time Updates**: Instant synchronization of note changes
- **Secure API**: RESTful API with proper authentication and validation
- **Docker Support**: Easy deployment with containerization

## 🛠 Tech Stack

### Backend
- **Go 1.23** - High-performance backend language
- **Fiber v2** - Lightweight web framework for Go
- **GORM** - ORM for database operations
- **PostgreSQL** - Robust relational database
- **JWT** - Secure token-based authentication
- **bcrypt** - Password hashing

### Frontend
- **React 18** - Modern JavaScript library for UI
- **TypeScript** - Type-safe JavaScript
- **Tailwind CSS** - Utility-first CSS framework
- **Axios** - HTTP client for API calls
- **React Router** - Client-side routing

### Infrastructure
- **Docker** - Containerization
- **Docker Compose** - Multi-container orchestration
- **Nginx** - Web server for frontend

## 📋 Prerequisites

Before running this application, make sure you have the following installed:

- **Docker** (version 20.10 or later)
- **Docker Compose** (version 2.0 or later)
- **Git** (for cloning the repository)

## 🏃‍♂️ Quick Start

1. **Clone the repository**
   ```bash
   git clone https://github.com/yourusername/fundament.git
   cd fundament
   ```

2. **Start the application**
   ```bash
   make start
   ```

3. **Access the application**
   - Frontend: http://localhost:3000
   - Backend API: http://localhost:8080
   - Database: localhost:5432

4. **Stop the application**
   ```bash
   make stop
   ```

## 📖 Usage

### Available Commands

```bash
make start    # Start the application (builds and runs all services)
make stop     # Stop the application and clean up containers
make help     # Show available commands
```

### First Time Setup

1. The application will automatically:
   - Build Docker images for backend and frontend
   - Start PostgreSQL database
   - Run database migrations
   - Start the backend API server
   - Serve the frontend via Nginx

2. Visit http://localhost:3000 to access the application

3. Register a new account or login with existing credentials

## 🔧 API Documentation

### Authentication Endpoints

#### Register User
```http
POST /api/auth/register
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123"
}
```

#### Login User
```http
POST /api/auth/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "password123"
}
```

### Notes Endpoints (Protected)

All notes endpoints require JWT authentication. Include the token in the Authorization header:
```
Authorization: Bearer <your-jwt-token>
```

#### Get All Notes
```http
GET /api/notes
```

#### Create Note
```http
POST /api/notes
Content-Type: application/json

{
  "content": "This is my note content"
}
```

#### Get Single Note
```http
GET /api/notes/:id
```

#### Update Note
```http
PUT /api/notes/:id
Content-Type: application/json

{
  "content": "Updated note content"
}
```

#### Delete Note
```http
DELETE /api/notes/:id
```

### Response Format

Success Response:
```json
{
  "notes": [
    {
      "id": 1,
      "user_id": 1,
      "content": "Note content",
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z"
    }
  ]
}
```

Error Response:
```json
{
  "error": "Error message description"
}
```

## 🏗 Project Structure

```
fundament/
├── backend/                    # Go backend application
│   ├── cmd/server/            # Application entry point
│   ├── internal/
│   │   ├── database/          # Database connection logic
│   │   ├── handlers/          # HTTP request handlers
│   │   ├── middleware/        # Authentication middleware
│   │   ├── models/           # Database models
│   │   └── utils/            # Utility functions
│   ├── Dockerfile            # Backend container config
│   └── go.mod               # Go module dependencies
├── frontend/                 # React frontend application
│   ├── src/
│   │   ├── components/       # Reusable UI components
│   │   ├── contexts/         # React context providers
│   │   ├── pages/           # Page components
│   │   ├── services/        # API service functions
│   │   ├── types/           # TypeScript type definitions
│   │   └── utils/           # Utility functions
│   ├── Dockerfile           # Frontend container config
│   ├── nginx.conf          # Nginx configuration
│   └── package.json        # Node.js dependencies
├── docker-compose.yml      # Multi-container orchestration
├── Makefile               # Development commands
└── README.md             # This file
```

## 🔐 Environment Variables

The application uses the following environment variables:

### Backend
- `DATABASE_URL` - PostgreSQL connection string
- `JWT_SECRET` - Secret key for JWT token signing
- `PORT` - Server port (default: 8080)
- `CORS_ORIGIN` - Allowed CORS origins

### Frontend
- `REACT_APP_API_URL` - Backend API URL

### Database
- `POSTGRES_USER` - PostgreSQL username
- `POSTGRES_PASSWORD` - PostgreSQL password
- `POSTGRES_DB` - PostgreSQL database name

## 🐳 Docker Services

The application consists of three main services:

1. **postgres** - PostgreSQL database (port 5432)
2. **backend** - Go API server (port 8080)
3. **frontend** - React application served via Nginx (port 3000)

## 🔒 Security Features

- **Password Hashing**: Uses bcrypt for secure password storage
- **JWT Authentication**: Stateless authentication with expiration
- **CORS Protection**: Configurable cross-origin resource sharing
- **Input Validation**: Server-side validation for all inputs
- **SQL Injection Protection**: Parameterized queries via GORM

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📝 Development

### Backend Development
```bash
cd backend
go mod download
go run cmd/server/main.go
```

### Frontend Development
```bash
cd frontend
npm install
npm start
```

### Database
The PostgreSQL database runs in a Docker container and automatically sets up the required tables using GORM migrations.

## 🐛 Troubleshooting

### Common Issues

1. **Port already in use**
   - Make sure ports 3000, 8080, and 5432 are available
   - Use `docker-compose down` to clean up previous containers

2. **Database connection errors**
   - Ensure PostgreSQL container is running
   - Check database credentials in docker-compose.yml

3. **Frontend not loading**
   - Verify backend is running on port 8080
   - Check CORS settings

### Logs
```bash
# View all service logs
docker-compose logs

# View specific service logs
docker-compose logs backend
docker-compose logs frontend
docker-compose logs postgres
```

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🙏 Acknowledgments

- [Fiber](https://gofiber.io/) - Web framework for Go
- [React](https://reactjs.org/) - Frontend library
- [Tailwind CSS](https://tailwindcss.com/) - CSS framework
- [PostgreSQL](https://postgresql.org/) - Database
- [Docker](https://docker.com/) - Containerization

---

**Happy note-taking! 📝**