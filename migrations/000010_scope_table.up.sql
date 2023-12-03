CREATE TABLE scope (
    id INT8 PRIMARY KEY DEFAULT unique_rowid(),
    name STRING(50) NOT NULL,
    description STRING NOT NULL
);