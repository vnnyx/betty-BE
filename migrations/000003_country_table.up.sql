CREATE TABLE country (
    id INT8 PRIMARY KEY DEFAULT unique_rowid(),
    iso STRING(10) NOT NULL UNIQUE,
    name STRING(100) NOT NULL UNIQUE
);