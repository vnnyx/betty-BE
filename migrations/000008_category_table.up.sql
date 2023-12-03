CREATE TABLE category (
    id INT8 PRIMARY KEY DEFAULT unique_rowid(),
    name STRING(100) NOT NULL,
    color STRING(255) NOT NULL,
    franchise_id INT8 REFERENCES franchise(id) ON DELETE CASCADE,
    UNIQUE(name, franchise_id)
);