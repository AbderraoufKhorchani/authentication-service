# Authentication Service

This Golang-based authentication service provides a secure and reliable solution for user authentication, utilizing PostgreSQL as the backend database. The service includes a well-documented API, powered by Swagger annotations, making it easy for developers to integrate authentication into their applications.

## Features

- User registration and login
- Secure password hashing and storage
- JWT token generation for authenticated users
- Swagger-documented API for easy integration and development

## Getting Started

### Prerequisites

To run this authentication service, ensure you have the following prerequisites installed:

- Go programming language
- PostgreSQL database 

### Installation

1. Clone the repository:

```bash
   git clone https://github.com/rf-krcn/authentication-service.git
   cd authentication-service
```
2. Install dependencies:
```go
   go install ./...
```
3. Configuration

Configure the service by setting up the PostgreSQL connection details by editing the DSN in the main.go file.

4. Run project
```go
   go run ./cmd
```


## Usage
Example API usage:
```bash
# Register a new user and obtain JWT token
curl -X POST http://localhost:8080/register -d '{"username": "exampleuser", "password": "securepassword"}'

# Login and obtain JWT token
curl -X POST http://localhost:8080/login -d '{"username": "exampleuser", "password": "securepassword"}'

# Reset password
curl -X POST http://localhost:8080/reset_password -d '{"user_name": "exampleuser", "password": "securepassword", "new_password": "newsecurepassword"}'
```
For more detailed information, refer to the [API Documentation](#api-doc) section.
## API Documentation<a id="api-doc"></a>

The API is thoroughly documented using Swagger annotations. Access the Swagger UI at [http://localhost:8080/docs/index.html](http://localhost:8080/docs/index.html) to explore and interact with the API.
