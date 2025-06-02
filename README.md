# Go Fiber Hexagonal API

This project is a RESTful API web server built using Go and the Fiber framework, following the hexagonal architecture pattern. It includes an authentication service and utilizes MongoDB for data storage, with GORM as the ORM.

## Project Structure

```
github.com/Gitong23/go-fiber-hex-api
├── cmd
│   └── main.go                # Entry point of the application
├── config
│   └── config.go              # Configuration settings
├── internal
│   ├── auth
│   │   ├── handler.go         # HTTP handlers for authentication
│   │   ├── service.go         # Business logic for authentication
│   │   └── repository.go      # Database interactions for authentication
│   ├── user
│   │   ├── handler.go         # HTTP handlers for user management
│   │   ├── service.go         # Business logic for user management
│   │   └── repository.go      # Database interactions for user data
│   ├── middleware
│   │   ├── auth.go            # Authentication middleware
│   │   └── logging.go         # Logging middleware
│   ├── response
│   │   └── response.go        # Global response structure
│   ├── errors
│   │   └── errors.go          # Custom error types and handling
│   └── logger
│       └── logger.go          # Logging configuration and functions
├── pkg
│   └── db
│       └── mongo.go           # MongoDB connection and GORM setup
├── docker-compose.yml          # Docker configuration for services
├── go.mod                      # Go module definition
├── go.sum                      # Dependency checksums
└── README.md                   # Project documentation
```

## Getting Started

### Prerequisites

- Go (version 1.16 or higher)
- Docker and Docker Compose

### Setup

1. Clone the repository:
   ```
   git clone <repository-url>
   cd github.com/Gitong23/go-fiber-hex-api
   ```

2. Build the Go application:
   ```
   go build -o app ./cmd/main.go
   ```

3. Start the MongoDB service using Docker Compose:
   ```
   docker-compose up -d
   ```

4. Run the application:
   ```
   ./app
   ```

### API Endpoints

- **Authentication**
  - `POST /auth/login` - Login user
  - `POST /auth/register` - Register new user

- **User Management**
  - `GET /user/:id` - Get user details
  - `PUT /user/:id` - Update user information

### Logging

The application includes logging middleware to log requests and responses for better debugging and monitoring.

### Error Handling

Custom error types and handling functions are implemented to manage application errors effectively.

### Contributing

Contributions are welcome! Please open an issue or submit a pull request for any improvements or bug fixes.

### License

This project is licensed under the MIT License. See the LICENSE file for details.