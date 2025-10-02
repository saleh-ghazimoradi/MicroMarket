CREATE TABLE IF NOT EXISTS order_catalog (
    order_id CHAR(27) REFERENCES orders (id) ON DELETE CASCADE,
    catalog_id CHAR(27),
    quantity INT NOT NULL,
    PRIMARY KEY (catalog_id, order_id)
);