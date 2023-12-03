CREATE TABLE attachment_file (
    id INT8 PRIMARY KEY DEFAULT unique_rowid(),
    path STRING(255) NOT NULL,
    is_deleted BOOL DEFAULT false
);