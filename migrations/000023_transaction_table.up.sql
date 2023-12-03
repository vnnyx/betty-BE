CREATE TABLE transaction (
    id INT8 PRIMARY KEY DEFAULT unique_rowid(),
    user_id INT8 REFERENCES users (id) ON DELETE CASCADE,
    variant_menu_id INT8 REFERENCES variant_menu (id) ON DELETE CASCADE,
    quantity INT NOT NULL,
    total_price INT8 NOT NULL
);

CREATE INDEX idx_user_variant_menu ON transaction (user_id, variant_menu_id);