-- +goose Up
-- +goose StatementBegin
CREATE TABLE "sales" (
    "id" uuid DEFAULT uuid_generate_v4() PRIMARY KEY,
    "product_id" uuid NOT NULL,
    "seller_id" uuid NOT NULL,
    "customer_name" varchar(255) NOT NULL,
    "total_itens" integer NOT NULL,
    "total_price" decimal NOT NULL,
    "is_pay" boolean DEFAULT false,
    "payed_at" timestamptz NOT NULL DEFAULT (now()),
    "created_at" timestamptz NOT NULL DEFAULT (now()),
    "updated_at" timestamptz NOT NULL DEFAULT (now())
)
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS sales;
-- +goose StatementEnd
