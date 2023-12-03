CREATE TABLE user_group (
    id INT8 PRIMARY KEY DEFAULT unique_rowid(),
    user_id INT8 REFERENCES users (id) ON DELETE CASCADE,
    role_id INT8 REFERENCES role (id) ON DELETE CASCADE
);

CREATE INDEX idx_user_role ON user_group (user_id, role_id);