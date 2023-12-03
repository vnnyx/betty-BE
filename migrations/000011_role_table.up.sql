CREATE TABLE role (
    id INT8 PRIMARY KEY DEFAULT unique_rowid(),
    name STRING(50) NOT NULL,
    color STRING(255) NOT NULL
);