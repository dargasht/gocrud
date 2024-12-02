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
