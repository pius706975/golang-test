# Backend Golang Test

## Table of Contents

- [Project Structure](#project-structure)
- [Installation](#installation)
- [Running the Application](#running-the-application)
- [API Documentation](#api-documentation)
- [Folder Structure](#folder-structure)
- [Contributor](#contributor)

---

## Project Structure

This project uses [Golang](https://golang.org/), [Gin](https://github.com/gin-gonic/gin) as the HTTP web framework, and [GORM](https://gorm.io/) for database ORM. 

## Installation

**1. Clone the repository:**
   ```bash
   git clone https://github.com/pius706975/golang-test.git
   ```

**2. Install dependencies:**
   ```bash
   go mod tidy
   ```

**3. Database configuration:**

   There are 2 options to configure the database. 

   - Using installed PostgreSQL on your computer
   - Using Docker Container

   I will explain for docker container to make it easier if you have installed docker container on your computer. 

  - Run postgres container using docker compose

    ```bash
        version: "3.8"

        services:
        postgres:
            image: postgres:17.0-alpine3.19
            container_name: local-postgres
            restart: always
            environment:
            POSTGRES_DB: db_test
            POSTGRES_USER: pius
            POSTGRES_PASSWORD: piuspius
            ports:
            - "5433:5432"
            volumes:
            - postgres_data:/var/lib/postgresql/data
            networks:
            - backend_network

        volumes:
        postgres_data:
            driver: local

        networks:
        backend_network:
            driver: bridge

      ```

    ***remove "sudo" if you don't use linux***
      ``` bash
      sudo docker compose up

      *Don't use "-d" if you want to see the process*
      ```

   - Check if the container is running 
      ``` bash
      sudo docker ps -a
      ```

   - If the container is running, execute the container.
      ``` bash
      sudo docker exec -it <CONTAINER ID> bash
      ```

   - Run the psql command 
      ```
      psql -U <POSTGRES_USER from docker compose> -d <POSTGRES_DB from docker compose>

      *e.g., "psql -U pius -d db_test"*

      ```

   - Copy the `.env.example` file to `.env`:
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

   - Update the `.env` file with your environment variables.

   - Migrate the database model into database
      ``` bash
        go run . migration 
        *or*
        go run . migration -u
      ```
  
  - Fill the tables with existing data
      ``` bash
        go run . seed
        *or*
        go run . seed -u
      ```  

  - Delete all data from database
      ``` bash
        go run . seed -d
      ```

  - Drop database
      ``` bash
        go run . migration -d
      ```

## Running the Application

To start the application, run:

```bash
go run . serve
```

To create a new transaction, we need to login first and use the token.

Use one of these accounts:
- email: jerry@gmail.com, tom@gmail.com
- password: User@123

## API Documentation

API documentation is generated using swaggo **v1.16.4**
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

## 👨‍💻 Author

- Pius Restiantoro - [GitHub](https://github.com/pius706975)