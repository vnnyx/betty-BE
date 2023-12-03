CREATE TABLE activity_field (
    id INT8 PRIMARY KEY DEFAULT unique_rowid(),
    name STRING(50) NOT NULL
);