# go-microservice-mongo
*Golang archetype oriented to microservices.*

## Dependencies
Golang, Docker, Make, [Swag tool](https://github.com/swaggo/swag)

## Features
- Architecture
    - Handlers, respositories and services
    - Custom Messages and Errors
    - Pagination and Ordering
- Go 1.23 (at the moment)
- Libraries
    - Web: Fiber
    - OAuth2: Gocloak
    - Validations: Go Playground Validator
    - Unit Test: Testify
    - DB: MongoDB
    - Tracing: Opentelemetry
    - Test: Testcontainers
    - OpenAPI: Fiber Swagger
- Keycloak as Auth Server
- Distributed tracing
    - OpenTelemetry, Micrometer and Jaeger
- Swagger
    - Swaggo & Fiber Swagger
    - Customized with command **make swagger** (OAuth2 server by parameter and not static)
- Custom Auditory
- Database
    - MongoDB for the app
    - Testcontainers for testing

## Files
- [Dockerfile](https://github.com/javiorfo/go-microservice-mongo/tree/master/Dockerfile)
- [Ship files](https://github.com/javiorfo/java-spring3-microservice/tree/master/ships)
    - For those using Neovim and [this plugin](https://github.com/javiorfo/nvim-ship)

## Usage
- Executing `make help` all the available commands will be listed. 
- Also the standard Go commands could be used, like `go run main.go`
- To use this archetype with a different name, execute this command to replace the names:
```bash
find . -type f -exec sed -i 's/go-microservice-mongo/your-project-name/g' {} +
```

## Postgres instead of MongoDB
- [Postgres repo](https://github.com/javiorfo/go-microservice) contains version with Postgres

---

### Donate
- **Bitcoin** [(QR)](https://raw.githubusercontent.com/javiorfo/img/master/crypto/bitcoin.png)  `1GqdJ63RDPE4eJKujHi166FAyigvHu5R7v`
