CREATE TABLE franchise (
    id INT8 PRIMARY KEY DEFAULT unique_rowid(),
    name STRING(100) NOT NULL UNIQUE,
    company_id INT8 REFERENCES company (id) ON DELETE CASCADE,
    photo_id INT8 DEFAULT NULL
);

CREATE INDEX idx_franchise_photo ON franchise (photo_id);