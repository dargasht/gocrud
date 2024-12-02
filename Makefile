sqlc: 
	sqlc vet
	sqlc generate

goose-up:
	goose -dir ./database/migration postgres $(app) up

goose-down:
	goose -dir ./database/migration postgres $(app) down

goose-status:
	goose -dir ./database/migration postgres $(app) status

goose-redo:
	goose -dir ./database/migration postgres $(app) redo

goose-reset:
	goose -dir ./database/migration postgres $(app) reset

build:
	@go build -o ./bin/main ./cmd/api/.  

run: build
	./bin/main

all: sqlc run