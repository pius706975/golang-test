# Golang Boilerplate with Gin and GORM
This is a simple Golang boilerplate with **Gin** framework with a ready-to-use configuration for backend development. You can adjust it according to your requirements.

## Table of Contents

- [Project Structure](#project-structure)
- [Installation](#installation)
- [Configuration](#configuration)
- [Running the Application](#running-the-application)
- [API Documentation](#api-documentation)
- [Folder Structure](#folder-structure)
- [Contributor](#contributor)

## Project Structure

This project uses [Golang](https://golang.org/), [Gin](https://github.com/gin-gonic/gin) as the HTTP web framework, and [GORM](https://gorm.io/) for database ORM. 

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/pius706975/golang-boilerplate-with-gin.git
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

## Configuration

1. Copy the `.env.example` file to `.env`:
   ```bash
    APP_PORT = 5000
    BASE_URL = http://localhost:5000/api
    MODE = development

    DB_PORT = <Db Port>
    DB_USERNAME = <Db Username>
    DB_PASSWORD = <Db Password>
    DB_NAME = <Db Name>
    DB_HOST = <Db Host>

    MAILER_PORT = <smtp port>
    MAILER_HOST = <smtp host>
    MAILER_EMAIL = <sender email>
    MAILER_PASSWORD = <password>

    JWT_ACCESS_TOKEN_SECRET = <Access Token Secret>
    JWT_REFRESH_TOKEN_SECRET = <Refresh Token Secret>
   ```

2. Update the `.env` file with your environment variables.

## Running the Application

To start the application, run:

```bash
go run . serve
```

To run database migration, use:
```bash
# migrate the database models
go run . migration -u 
# drop database
go run . migration -d
```
## API Documentation

API documentation is generated using Swagger v1.16.4
. You can access the documentation by running the server and visiting `<your base url>/docs/index.html` in your browser.

### Generating Swagger Docs

To update Swagger documentation, run:
```bash
swag init
```
Make sure you have installed swaggo globally on your computer.
Read the swaggo documentation [here](https://pkg.go.dev/github.com/swaggo/swag/v2#readme-getting-started)

## Folder Structure

Here's a breakdown of the project folder structure:

- **api/**: Handles route definitions and server setup
  - **routes/**: Defines all API routes from modules
  - **server.go**: Configures and initializes the server

- **cmd/**: Contains command line scripts
  - **command.line.go**: CLI execution entry point

- **config/**: Configuration files, including environment variables
  - **env.go**: Loads and parses environment variables

- **docs/**: API documentation files (Swagger)
  - **swagger.json** and **swagger.yaml**: Swagger specification files

- **interfaces/**: Interfaces for abstracting logic
  - **auth.interface.go** and **user.interface.go**: Define interface contracts for auth and user modules

- **middlewares/**: Middleware functions for request handling
  - **auth.middleware.go**: Authorization middleware
  - **jwt.service.go**: JWT utility functions

- **modules/**: Core application modules
  - **auth/**: Authentication module
  - **user/**: User-related functionality
  - **other module/**

- **package/**: Reusable packages
  - **database/**: Database configuration, models, and migrations
    - **models/**: GORM models
    - **config.go**: Database connection configuration
    - **migrations.go**: Database migration logic
  - **utils/**: Utility functions

- **main.go**: Application entry point

- **.env.example**: Example environment configuration

## 👨‍💻 Contributor

- Pius Restiantoro - [GitHub](https://github.com/pius706975)