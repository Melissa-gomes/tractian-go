-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

CREATE TABLE "products" (
    "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    "name" varchar(255) NOT NULL,
    "price" decimal NOT NULL,
    "quantity" integer NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
)

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP TABLE IF EXISTS products;
-- +goose StatementEnd
