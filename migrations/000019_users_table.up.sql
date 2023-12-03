CREATE TABLE users (
    id INT8 PRIMARY KEY DEFAULT unique_rowid(),
    is_super_admin BOOL NOT NULL,
    is_admin BOOL NOT NULL,
    email STRING(320) NOT NULL UNIQUE,
    password STRING(255) DEFAULT NULL,
    fullname STRING(255) NOT NULL,
    phone_number STRING(50) DEFAULT NULL,
    company_id INT8 REFERENCES company (id) ON DELETE CASCADE,
    franchise_id INT8 REFERENCES franchise (id) ON DELETE CASCADE,
    is_active BOOL NOT NULL DEFAULT TRUE,
    photo_id INT8 REFERENCES attachment_file (id) ON DELETE SET NULL,
    refresh_token STRING DEFAULT NULL,
    shared_secret STRING DEFAULT NULL
);