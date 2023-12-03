CREATE TABLE company (
    id INT8 PRIMARY KEY DEFAULT unique_rowid(),
    brand_name STRING(100) UNIQUE NOT NULL,
    franchise_name STRING(100) NOT NULL,
    address_1 STRING(250) DEFAULT NULL,
    address_2 STRING(250) DEFAULT NULL,
    city_id INT8 REFERENCES city (id) ON DELETE CASCADE,
    country_id INT8 REFERENCES country (id) ON DELETE CASCADE,
    photo_id INT8 REFERENCES attachment_file (id) ON DELETE SET NULL,
    postal_code STRING(10) DEFAULT NULL
);