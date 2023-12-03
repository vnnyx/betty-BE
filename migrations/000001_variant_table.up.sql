CREATE TABLE variant (
    id INT8 PRIMARY KEY DEFAULT unique_rowid(),
    name STRING(100) NOT NULL,
    is_multi BOOL DEFAULT false,
    is_deleted BOOL DEFAULT false
);