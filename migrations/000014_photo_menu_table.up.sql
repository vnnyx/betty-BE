CREATE TABLE photo_menu (
    id INT8 PRIMARY KEY DEFAULT unique_rowid(),
    menu_id INT8 REFERENCES menu (id) ON DELETE CASCADE,
    attachment_file_id INT8 REFERENCES attachment_file (id) ON DELETE CASCADE,
    is_deleted BOOL DEFAULT false
);