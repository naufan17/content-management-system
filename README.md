# Content Management System
A simple content management system API for managing news

## Table of Content
- [Run Server Locally](#run-server-locally)
- [Run Server with Docker Compose](#run-server-with-docker-compose)
- [Feature](#feature)
- [Project Structure](#project-structure)
- [API Documentation](#api-documentation)
- [Response Format](#response-format)

## Run Server Locally
- Run server
```
go run cmd/server/main.go
```

## Run Server with Docker Compose
- Run server
```
docker compose up --build
```
- Stop server
```
docker compose down
```

## Feature
- **Database**: using https://gorm.io/driver/postgres
- **ORM**: using https://gorm.io/
- **Authentication**: using https://github.com/dgrijalva/jwt-go
- **Validation**: using https://github.com/go-playground/validator/v10
- **Rate Limit**: using https://golang.org/x/time/rate
- **Security**: https://github.com/danielkov/gin-helmet
- **CORS**: using https://github.com/gin-contrib/cors
- **Environtment variables**: using https://github.com/joho/godotenv
- **API documentation**: using https://github.com/swaggo/swag, https://github.com/swaggo/gin-swagger and https://github.com/swaggo/files

## Project Structure
```
cmd\
  |--server\       # Command to run server
  |--database\     # Command database migration and seeder          
config\            # Configuration
database\
  |--seeders\      # Database seeder
internal\
  |--dtos\         # Data transfer object
  |--handlers\     # Request handlers
  |--middlewares\  # Middleware
  |--models\       # Database models
  |--repositories\ # Database queries
  |--services\     # Business logic
pkg\
  |--auth\         # Authentication
  |--util\         # Utility function
route\             # API routes
```

## API Documentation
To view the API documentation, open the following link:
<br/>
``GET /swagger/index.html`` - View API documentation

## Schema Database
https://www.dbdiagram.io/d/CMS-6676cf305a764b3c7223dcee

## Response Format
- **Success Response**:
```
{
  "data": {
    "key": "value",
  }
}
```
OR
```
{
  "message": "success message",
}
```
- **Error Response**:
```
{
  "error": "error message",
}
```