CREATE TABLE menu_ingredient (
    id INT8 PRIMARY KEY DEFAULT unique_rowid(),
    ingredient_id INT8 REFERENCES ingredient (id) ON DELETE CASCADE,
    menu_id INT8 REFERENCES menu (id) ON DELETE CASCADE,
    quantity FLOAT8 NOT NULL
);

CREATE INDEX idx_menu_ingredient ON menu_ingredient (ingredient_id, menu_id);