CREATE TABLE menu (
    id INT8 PRIMARY KEY DEFAULT unique_rowid(),
    matchcode STRING(50) NOT NULL UNIQUE,
    name STRING(100) NOT NULL,
    stock INT8 DEFAULT 0,
    price FLOAT8 NOT NULL,
    description STRING NOT NULL,
    status_id INT DEFAULT NULL,
    franchise_id INT8 REFERENCES franchise (id) ON DELETE CASCADE,
    popularity_count INT8 DEFAULT 0,
    is_deleted BOOL DEFAULT false
);

CREATE INDEX idx_menu_status ON menu (status_id);