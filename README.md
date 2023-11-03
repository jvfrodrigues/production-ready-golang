# WEX Transaction Product Test

Project given by WEX as code challenge, a API that stores transaction and allows users to retrieve those transactions converting the value of the trasaction to the speciefied country's currency

## Language

This project is written in _Golang_ 1.21

## Important libraries

### [Gin](https://github.com/gin-gonic/gin)

Gin is our HTTP Server, it handles our routing system and HTTP requests

### [Swaggo](https://github.com/swaggo/swag)

Swag converts Go annotations to Swagger Documentation 2.0.

### [Gorm](https://github.com/go-gorm/gorm)

ORM library for dealing with database connections

### [Validator](https://github.com/asaskevich/govalidator)

Helps in validating our entities and dtos

### [Zap](https://pkg.go.dev/go.uber.org/zap)

Leveled loggin library

### [Testify](github.com/stretchr/testify)

Testing tools

## How to run

The project can run through Docker or locally first will need to

1. Download and install [Go](https://go.dev/) and/or [Docker](https://www.docker.com/)
2. Download this repository

   ```bash
    git clone https://github.com/jvfrodrigues/transaction-product-wex
   ```

3. Create .env file copying .env.example

   ```bash
   cp .env.example .env
   ```

- Using Docker

4. Run docker compose

   ```bash
    docker-compose -f docker-compose-prod.yaml up -d
   ```

5. Test! It should run on [localhost:8080](http://localhost:8080)

6. If you want you can set the env variable in .env to "prd" so it will use postgres instead of sqlite

- Locally

4. Install all dependencies
   ```bash
   go mod download
   ```
5. Run the project
   ```bash
   go run main.go
   ```
6. Test! It should run on [localhost:8080](http://localhost:8080)
