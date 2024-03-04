# Feature

See [feature](FEATURE.md)

# Structure

```
├── Makefile // Scripts to perform various build, install, analysis, etc operations.
├── cmd // Main applications
│   ├── main.go
│   └── wire // auto gen DI
│       ├── app.go // app interface
│       ├── wire.go
│       └── wire_gen.go
├── config
│   ├── config.go // require environment variables
│   └── constants.go // default environment variables
├── constants // constants variables
├── domain
│   ├── entity // model interface
│   └── repository // repository interface
├── handler // controller, handle request
├── infra
│   ├── elasticsearch // elasticsearch adapter
│   ├── mongodb // mongodb adapter
│   ├── postgresql
│   │   ├── database.go
│   │   ├── logger.go
│   │   ├── migrate.go // atlas adapter
│   │   ├── migrations // version migrations
│   │   ├── model // database model
│   │   └── repository // database repository
│   ├── redis // redis adapter
│   └── s3 // aws s3 adapter
├── infrastructure_config // Config for docker infra
├── logs // logs folder
├── middlewares // put the before & after logic of handle request
├── mock // mock for unit testing
├── pkg // utils for service
│   └── errors
│       └── custom
│           ├── internal_server.go // 5xx error
│           ├── logic.go // logic error
│           ├── main.go
│           └── validate.go // validation error
├── routers // router for services use REST API
├── scripts // scripts folder
├── template // contain fixed template for email, presentation,...
├── usecase // business logic
└── validations // custom validation
```

# Apply git hook

```shell
$ make hook
```

# Generate wire Dependencies Injection

```shell
$ make gen_di
```

# Generate module

Name convention: snake_case (ex: order_product)

```shell
$ make gen module=order
```

# Mock for testing

Name convention: snake_case (ex: order_product)

## Repository

```shell
$ make gen_mock repository=product
```

## UseCase

```shell
$ make gen_mock usecase=product
```

## Service

```shell
$ make gen_mock service=data
```

# Migration

## Re-hash atlas

```shell
$ atlas migrate hash --dir file://infra/postgresql/migrations
```
