CREATE TABLE menu_category (
    id INT8 PRIMARY KEY DEFAULT unique_rowid(),
    menu_id INT8 REFERENCES menu (id) ON DELETE CASCADE,
    category_id INT8 REFERENCES category (id) ON DELETE CASCADE
);

CREATE INDEX idx_menu_category ON menu_category (menu_id, category_id);