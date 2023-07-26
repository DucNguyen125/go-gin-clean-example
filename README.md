# Structure

```
├── api // OpenAPI/Swagger specs, JSON schema files protocol definition files.
├── cmd // Main applications
├── config // Contain configuration
│ ├── config.go // environment variables
│ └── constants.go // constants variables
├── docker-compose.yml // Container configuration
├── Dockerfile // Define container image
├── domain // Domain layer, interface for database layer
│ ├── entity // model interface
│ └── repository // repository interface
├── go.mod
├── go.sum
├── handler // controller, handle request
├── infra // Database layer
│ ├── postgresql
│ ├── database.go
│ ├── logger.go
│ ├── migrate.go
│ ├── model // database model
│ └── repository // database repository
├── logs // Contains log file
├── Makefile // Scripts to perform various build, install, analysis, etc operations.
├── middlewares // put the before & after logic of handle request
├── pkg // utils for service
├── README.md
├── routers // router for services use REST API
├── scripts
└── usecase // business logic
```

# Start Server

## Local

```shell
1. Copy file .env
$ cp .env.example .env

2. Run database
$ docker-compose up -d --no-deps database

3. Run server local
$ go run cmd/main.go
```

## With docker

```shell
1. Run all containers
$ docker-compose up -d
```

# Mock for testing

Generate mock package `ProductRepository` example

## Repository

```shell
$ mockgen -package=repository base-gin-golang/domain/repository ProductRepository > mock/domain/repository/product.go
```

## UseCase

```shell
$ mockgen -package=helpdesk_contract base-gin-golang/usecase/helpdesk_contract HelpDeskContractUseCase > mock/usecase/helpdesk_contract/helpdesk_contract.go
```

#Generate Html Swagger Docs

```shell
$ redoc-cli build api/docs/openapi.yaml
```
