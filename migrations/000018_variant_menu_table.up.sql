CREATE TABLE variant_menu (
    id INT8 PRIMARY KEY DEFAULT unique_rowid(),
    variant_id INT8 REFERENCES variant (id) ON DELETE CASCADE,
    menu_id INT8 REFERENCES menu (id) ON DELETE CASCADE,
    price INT8 NOT NULL,
    is_deleted BOOL DEFAULT false
);

CREATE INDEX idx_variant_menu ON variant_menu (variant_id, menu_id);