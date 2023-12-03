CREATE TABLE unit (
    id INT8 PRIMARY KEY DEFAULT unique_rowid(),
    name STRING(100) NOT NULL UNIQUE,
    is_international BOOL DEFAULT false,
    conversion_rate FLOAT8 DEFAULT NULL,
    conversion_id INT8 DEFAULT NULL,
    franchise_id INT8 REFERENCES franchise (id) ON DELETE CASCADE,
    is_deleted BOOL DEFAULT false,
    UNIQUE (name, franchise_id)
);

CREATE INDEX idx_unit_conversion ON unit (conversion_id);