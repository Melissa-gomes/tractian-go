-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';

ALTER TABLE
    sales
ADD
    CONSTRAINT fk_product_id_in_sales FOREIGN KEY (product_id) REFERENCES products (id)


ALTER TABLE
    sales
ADD
    CONSTRAINT fk_seller_id_in_sales FOREIGN KEY (seller_id) REFERENCES sellers (id)
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

ALTER TABLE
    sales DROP CONSTRAINT fk_product_id_in_sales

ALTER TABLE
    sales DROP CONSTRAINT fk_seller_id_in_sales
-- +goose StatementEnd
