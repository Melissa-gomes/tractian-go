-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE "sellers" (
    "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    "name" varchar(255) NOT NULL,
    "is_active" boolean DEFAULT false,
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
)

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP TABLE IF EXISTS sellers;
-- +goose StatementEnd
