# Simple Web Server

A Go-based web server application with PostgreSQL database integration.

## Project Structure

```
.
├── cmd/
│   └── main.go         # Application entry point
├── internal/
│   ├── handlers/       # HTTP request handlers
│   ├── models/         # Data models and database schemas
│   ├── services/       # Business logic and service layer
│   └── helper/         # Utility functions and helpers
├── go.mod             # Go module definition
├── go.sum             # Go module checksums
└── .env               # Environment variables (not tracked in git)
```

## Prerequisites

- Go 1.21.8 or higher
- PostgreSQL database
- Git

## Dependencies

- `github.com/jackc/pgx/v5` - PostgreSQL driver
- `github.com/joho/godotenv` - Environment variable management

## Setup Instructions

1. Clone the repository:
   ```bash
   git clone https://github.com/RaihanurRahman2022/simple-web-server.git
   cd simple-web-server
   ```

2. Install dependencies:
   ```bash
   go mod download
   ```

3. Create a `.env` file in the root directory with the following variables:
   ```
   DB_HOST=localhost
   DB_PORT=5432
   DB_USER=your_username
   DB_PASSWORD=your_password
   DB_NAME=your_database_name
   ```

4. Run the application:
   ```bash
   go run cmd/main.go
   ```

## Development Guidelines

### Code Organization

1. **Handlers (`internal/handlers/`)**
   - Handle HTTP requests and responses
   - Validate input data
   - Call appropriate services
   - Return HTTP responses

2. **Models (`internal/models/`)**
   - Define data structures
   - Include database schema definitions
   - Implement data validation methods

3. **Services (`internal/services/`)**
   - Implement business logic
   - Handle database operations
   - Process data transformations

4. **Helpers (`internal/helper/`)**
   - Provide utility functions
   - Implement common operations
   - Handle error formatting

### Best Practices

1. **Error Handling**
   - Use proper error wrapping
   - Implement meaningful error messages
   - Handle database errors appropriately

2. **Database Operations**
   - Use prepared statements
   - Implement proper connection pooling
   - Handle transactions appropriately

3. **Code Style**
   - Follow Go standard formatting
   - Write meaningful comments
   - Use meaningful variable names

4. **Testing**
   - Write unit tests for services
   - Implement integration tests
   - Test error scenarios

## Contributing

1. Create a new branch for your feature
2. Make your changes
3. Write tests if applicable
4. Submit a pull request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
