# Gin REST API

This project is a REST API built with the Gin framework in Go. It provides user and event management functionalities, including user registration, login, event creation, updating, deletion, and user registration for events.

## Features

- **User Management**:
  - User Signup
  - User Login
  - Password Hashing and Verification

- **Event Management**:
  - Create Event
  - Update Event
  - Delete Event
  - Register for Event

- **Database**:
  - SQLite database for storing user and event data
  - Write-Ahead Logging (WAL) mode for better concurrency

- **Security**:
  - Password hashing using Argon2
  - Authentication middleware for protected routes

## Getting Started

### Prerequisites

- Go 1.16 or higher
- SQLite3

### Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/mdombrov-33/gin-rest-api.git
   cd gin-rest-api
   ```

 2. Install dependencies:
    
    ```sh
    go mod tidy
    ```
  
3. Build the application:

    ```sh
    go build -o gin-rest-api
    ```

4. Run the application:

   ```sh
   ./gin-rest-api
   ```

  ## API Endpoints
  
### User Management
- **Signup: POST /signup**:

Request body:

 ```sh
 {
 "email": "testuser@gmail.com",
 "password": "test"
 }
```

Response: 201 Created

```sh
{
  "message": "User Created!"
}
```
