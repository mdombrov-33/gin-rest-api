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

- **Login: POST /login**:

Request body:

```sh
{
   "email": "testuser@gmail.com",
   "password": "test"
}
```

Response: 200 OK

```sh
{
  "message": "Login successful!",
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3R1c2VyQGdtYWlsLmNvbSIsImV4cCI6MTczMDE4NjYzMSwidXNlcklkIjo1fQ.aEZHPT_HdpHf6S3I-oIkgfHLYHUYFrBvXFlQQDaMMVA"
}
```

### Event Management

- **Create Event: POST /events**

Request body:
```sh
authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3R1c2VyQGdtYWlsLmNvbSIsImV4cCI6MTczMDE4NjYzMSwidXNlcklkIjo1fQ.aEZHPT_HdpHf6S3I-oIkgfHLYHUYFrBvXFlQQDaMMVA

{
"title": "Event Title",
"description": "Event Description",
"location": "Event Location",
"dateTime": "2025-12-31T23:59:59Z"
}
```

Response: 201 Created

```sh
{
"event": {
"ID": 1,
"Title": "Event Title",
"Description": "Event Description",
"Location": "Event Location",
"DateTime": "2025-12-31T23:59:59Z",
"UserID": 1
},
"message": "Event Created!"
}
```

- **Get All Events: GET /events**

Response: 200 OK

```sh
[
{
"ID": 1,
"Title": "Event Title",
"Description": "Event Description",
"Location": "Event Location",
"DateTime": "2025-12-31T23:59:59Z",
"UserID": 1
  }
]
```

- **Get Single Event: GET /events/:id**

Response: 200 OK

```sh
{
"ID": 1,
"Title": "Event Title",
"Description": "Event Description",
"Location": "Event Location",
"DateTime": "2025-12-31T23:59:59Z",
"UserID": 1
}
```

- **Update Event: PUT /events/:id**

Request Body:
```sh
authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3R1c2VyQGdtYWlsLmNvbSIsImV4cCI6MTczMDE4NjYzMSwidXNlcklkIjo1fQ.aEZHPT_HdpHf6S3I-oIkgfHLYHUYFrBvXFlQQDaMMVA

{
"title": "New Event Title",
"description": "New Event Description",
"location": "New Event Location",
"dateTime": "2065-12-31T23:59:59Z"
}
```

Response: 200 OK

```sh
{
"event": {
"ID": 1,
"Title": "New Event Title",
"Description": "New Event Description",
"Location": "New Event Location",
"DateTime": "2065-12-31T23:59:59Z",
"UserID": 1
},
"message": "Event updated!"
}
```

- **Delete Event: DELETE /events/:id**

Request Body:
```sh
authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3R1c2VyQGdtYWlsLmNvbSIsImV4cCI6MTczMDE4NjYzMSwidXNlcklkIjo1fQ.aEZHPT_HdpHf6S3I-oIkgfHLYHUYFrBvXFlQQDaMMVA
```

Response: 200 OK

```sh
{
"event": {
"ID": 1,
"Title": "Event Title",
"Description": "Event Description",
"Location": "Event Location",
"DateTime": "2025-12-31T23:59:59Z",
"UserID": 1
},
"message": "Event deleted!"
}
```

- **Register for Event: POST /events/:id/register**

Request Body:
```sh
authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3R1c2VyQGdtYWlsLmNvbSIsImV4cCI6MTczMDE4NjYzMSwidXNlcklkIjo1fQ.aEZHPT_HdpHf6S3I-oIkgfHLYHUYFrBvXFlQQDaMMVA
```

Response: 201 Created

```sh
{
"message": "Successfully registered for the event"
}
```

- **Cancel Registration for Event: DELETE /events/:id/register**

Request body:
```sh
authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3R1c2VyQGdtYWlsLmNvbSIsImV4cCI6MTczMDE4NjYzMSwidXNlcklkIjo1fQ.aEZHPT_HdpHf6S3I-oIkgfHLYHUYFrBvXFlQQDaMMVA
 ```

Response:
```sh
{
"message": "Successfully canceled registration"
}
```
