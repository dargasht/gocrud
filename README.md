# gocrud

Template for go crud applications we use. Also some generic handler functions and helpers we use.

## Installation

```bash
go get github.com/lestrrat-go/gocrud
```

## What we use

- [go-fiber](https://github.com/gofiber/fiber)
 for http server
- [pgx/v5](https://github.com/jackc/pgx)
 for database
- [pgxpool](https://github.com/jackc/pgx-pool)
 for database connection pool
- [sonic](https://github.com/bytedance/sonic)
 for json marshall/unmarshall
- [zap](https://github.com/uber-go/zap)
 for logging
- [sqlc](https://github.com/sqlc-dev/sqlc)
 for database code generation
- [sqlc](https://github.com/jmoiron/sqlx)
 for some wierd queries that can't be done with sqlc
- [goose](https://github.com/pressly/goose)
 for database migrations

## what we enjoy

- excessive use of postgres json_agg function
- simple architecture with no magic (maybe some magic, but not too much)
- no ORM
- writing raw sql queries
- simplicity simplicity and simplicity

## Folder structure

```bash
.
├── bin # for binaries (should be added in gitignore)
├── cfg # for configs
├── cmd
│   └── api
│       └── main.go # the main entry point
├── database
│   ├── migration # for database migrations
│   ├── query # for database queries
│   ├── repo # for sqlc generated queries and models 
│   └── store # for specific database implementations that uses sqlx
├── handler # for http handlers (data transformation and validation)
├── middleware # for http middlewares (if needed)
├── model # for service models and handler models (DTOs)
├── router # for instantiating handlers and injecting dependencies
├── service # for business logic (if needed)
└── utils # for utility functions 

```

Things to keep in mind:

- services should be stateless and not depend on external services
- use service layer if business logic is not simple otherwise use handler layer
- go to each folder in the internal folder and read the README.md

Usefull stuff in this package are:

- [error handler](https://github.com/dargasht/gocrud/blob/main/error_handler.go)

- [error helper stuff](https://github.com/dargasht/gocrud/blob/main/error_helper.go)
  - This contains stuff usefull to use in your handlers, they are (Just read comments in code):
  - [Standard error type we use](https://github.com/dargasht/gocrud/blob/main/error_helper.go#L32)
  - [some functions for creating errors](https://github.com/dargasht/gocrud/blob/main/error_helper.go#L45)
  - [NewNotFoundError](https://github.com/dargasht/gocrud/blob/main/error_helper.go#L58)
  - [Database Errors](https://github.com/dargasht/gocrud/blob/main/error_helper.go#L69)
  - [JSON and Validation Errors](https://github.com/dargasht/gocrud/blob/main/error_helper.go#L139)
  - [Token and permission errors](https://github.com/dargasht/gocrud/blob/main/error_helper.go#L162)

- [handler helper stuff](https://github.com/dargasht/gocrud/blob/main/handler_helper.go)
  - This contains stuff usefull to use in your handlers, they are (Just read comments in code):
  - [Handler config example](https://github.com/dargasht/gocrud/blob/main/handler_helper.go#L13)
  - [Standard responses](https://github.com/dargasht/gocrud/blob/main/handler_helper.go#L32)
  - [3 functions for pagination](https://github.com/dargasht/gocrud/blob/main/handler_helper.go#L78)
  - [EnsureAdmin function](https://github.com/dargasht/gocrud/blob/main/handler_helper.go#L145)
  - [Authenticate function](https://github.com/dargasht/gocrud/blob/main/handler_helper.go#L161)
  - [GetUserIDFromJWT function](https://github.com/dargasht/gocrud/blob/main/handler_helper.go#L190)
  - [EnsureAdmin function](https://github.com/dargasht/gocrud/blob/main/handler_helper.go#L145)
