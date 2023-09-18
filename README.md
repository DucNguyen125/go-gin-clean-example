# Structure

```
├── api // OpenAPI/Swagger specs, JSON schema files protocol definition files.
├── cmd // Main applications
├── config // Contain configuration
│   ├── config.go // environment variables
│   └── constants.go // constants variables
├── docker-compose.yml // Container configuration
├── Dockerfile // Define container image
├── domain // Domain layer, interface for database layer
│   ├── entity // model interface
│   └── repository // repository interface
├── error // error constants
├── go.mod
├── go.sum
├── handler // controller, handle request
├── infra // Database layer
│   └── postgresql
│       ├── database.go
│       ├── logger.go
│       ├── migrate.go
│       ├── model // database model
│       └── repository // database repository
├── infrastructure_config // Config for docker infra
├── logs // Contains log file
├── Makefile // Scripts to perform various build, install, analysis, etc operations.
├── middlewares // put the before & after logic of handle request
├── mock // for unit testing
├── pkg // utils for service
├── README.md
├── routers // router for services use REST API
├── template // contain fixed template for email, presentation,...
├── scripts
├── usecase // business logic
└── validations // custom validation
```

# Mock for testing

Generate mock package `ProductRepository` example

## Repository

```shell
$ mockgen -package=repository base-gin-golang/domain/repository ProductRepository > mock/domain/repository/product.go
```

## UseCase

```shell
$ mockgen -package=product base-gin-golang/usecase/product UseCase > mock/usecase/product/main.go
```

## Service

```shell
$ mockgen -package=errors base-gin-golang/pkg/errors Service > mock/pkg/errors/errors.go
```

# Generate Html Swagger Docs

```shell
$ redoc-cli build api/docs/openapi.yaml
```

# Apply git hook

```shell
make hook
```

# Generate wire Dependencies Injection

```shell
cd cmd/wire
wire
```

