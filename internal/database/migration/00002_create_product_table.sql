-- +goose Up
-- +goose StatementBegin
CREATE TABLE product (
	id SERIAL PRIMARY KEY,
	name VARCHAR(255) NOT NULL,
	price BIGINT NOT NULL,
	created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
DROP TABLE product;

-- +goose StatementEnd