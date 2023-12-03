CREATE TABLE ingredient (
    id INT8 PRIMARY KEY DEFAULT unique_rowid(),
    name STRING(100) NOT NULL,
    minimum_stock FLOAT8 NOT NULL,
    recent_stock FLOAT8 NOT NULL,
    price FLOAT8 NOT NULL,
    unit_id INT8 REFERENCES unit (id) ON DELETE CASCADE,
    franchise_id INT REFERENCES franchise (id) ON DELETE CASCADE,
    is_deleted BOOL DEFAULT false
);