# devsMailGo API

Enterprise-standard REST API for devsMailGo email server management, built with Go.

## 🚀 Features

- **Enterprise Security**: JWT authentication, role-based access control, rate limiting
- **Comprehensive API**: Full CRUD operations for users, domains, aliases, mailing lists
- **Production Ready**: Structured logging, configuration management, health checks
- **Scalable Architecture**: Clean separation of concerns with service-oriented design
- **API Documentation**: Complete Postman collection with authentication examples

## 📋 Prerequisites

- Go 1.23+
- MySQL 8.0+
- Git

## 🛠️ Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd devsMailGo
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Configure environment**
   ```bash
   cp env.example .env
   # Edit .env with your database and configuration settings
   ```

4. **Run database migrations**
   ```bash
   go run cmd/main.go
   ```

## 🏗️ Project Structure

```
devsMailGo/
├── api/dto/           # Data Transfer Objects
├── cmd/              # Application entry point
├── config/           # Configuration management
├── controller/       # HTTP request handlers
├── middleware/       # HTTP middleware (auth, CORS, rate limiting)
├── models/           # Database models
├── repository/       # Data access layer
├── routes/           # Route definitions
├── service/          # Business logic layer
├── tests/            # Integration tests
├── utils/            # Utility functions
├── logger/           # Logging system
├── .env              # Environment variables
├── env.example       # Environment template
├── go.mod            # Go module file
├── go.sum            # Go dependencies checksum
└── README.md         # This file
```

## 🔧 Configuration

### Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `DB_HOST` | Database host | localhost |
| `DB_PORT` | Database port | 3306 |
| `DB_USER` | Database user | devsmailgo |
| `DB_PASSWORD` | Database password | (required) |
| `DB_NAME` | Database name | devsmailgo |
| `JWT_SECRET` | JWT signing secret | (change in production) |
| `JWT_EXPIRY_HOURS` | JWT token expiry | 24 |
| `SERVER_PORT` | Server port | 8080 |
| `ENVIRONMENT` | Environment (dev/prod) | development |
| `LOG_LEVEL` | Logging level | info |
| `RATE_LIMIT_REQUESTS` | Rate limit requests per window | 100 |
| `RATE_LIMIT_WINDOW` | Rate limit window | 1m |

## 🔐 Authentication

### Login
```bash
POST /api/login
Content-Type: application/json

{
  "username": "admin@domain.com",
  "password": "yourpassword"
}
```

### Using Authentication
Add the JWT token to all protected endpoints:
```bash
Authorization: Bearer <your-jwt-token>
```

## 📚 API Endpoints

### Public Endpoints
- `GET /api/health` - Health check
- `POST /api/login` - User authentication

### Protected Endpoints

#### Users
- `GET /api/users` - List all users
- `GET /api/users/{id}` - Get user by ID
- `POST /api/users` - Create user
- `PUT /api/users/{id}` - Update user
- `DELETE /api/users/{id}` - Delete user

#### Domains
- `GET /api/domains` - List all domains
- `GET /api/domain/{domain}` - Get domain
- `POST /api/domain/{domain}` - Create domain
- `PUT /api/domain/{domain}` - Update domain
- `DELETE /api/domain/{domain}` - Delete domain

#### Security
- `GET /api/banned` - List banned IPs
- `POST /api/banned/unban` - Unban IP
- `GET /api/jails` - List jails

#### Email Management
- `GET /api/aliases` - List aliases
- `POST /api/alias/{address}` - Create alias
- `GET /api/mailing-lists` - List mailing lists
- `POST /api/mailing-list/{address}` - Create mailing list

#### Monitoring
- `GET /api/logs` - List logs
- `GET /api/quota` - List quota
- `GET /api/roundcube-users` - List Roundcube users

## 🧪 Testing

Run integration tests:
```bash
go test ./tests/...
```

Run all tests with coverage:
```bash
go test -v -cover ./...
```

## 🚀 Deployment

### Development
```bash
go run cmd/main.go
```

### Production
```bash
# Build the application
go build -o bin/devsmailgo-api cmd/main.go

# Run the binary
./bin/devsmailgo-api
```

### Docker (Recommended)
```dockerfile
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o devsmailgo-api cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /app/devsmailgo-api .
EXPOSE 8080
CMD ["./devsmailgo-api"]
```

## 📊 Monitoring

### Health Check
```bash
curl http://localhost:8080/api/health
```

### Logs
The application uses structured logging with configurable levels:
- `DEBUG` - Detailed debug information
- `INFO` - General information
- `WARN` - Warning messages
- `ERROR` - Error messages
- `FATAL` - Fatal errors

## 🔒 Security Features

- **JWT Authentication**: Secure token-based authentication
- **Role-Based Access Control**: Different permission levels
- **Rate Limiting**: Protection against abuse
- **CORS Protection**: Cross-origin request handling
- **Security Headers**: XSS, CSRF protection
- **Input Validation**: Request payload validation

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🆘 Support

For support, email support@example.com or create an issue in the repository.

## 🔄 Version History

- **v1.0.0** - Initial release with enterprise features
- **v1.1.0** - Added rate limiting and enhanced security
- **v1.2.0** - Improved logging and monitoring 