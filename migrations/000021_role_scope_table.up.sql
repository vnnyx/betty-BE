CREATE TABLE role_scope (
    id INT8 PRIMARY KEY DEFAULT unique_rowid(),
    role_id INT8 REFERENCES role (id) ON DELETE CASCADE,
    scope_id INT8 REFERENCES scope (id) ON DELETE CASCADE
);

CREATE INDEX idx_role_scope ON role_scope (role_id, scope_id);