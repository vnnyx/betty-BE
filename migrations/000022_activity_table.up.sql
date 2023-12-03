CREATE TABLE activity (
    id INT8 PRIMARY KEY DEFAULT unique_rowid(),
    activity_field_id INT8 REFERENCES activity_field (id) ON DELETE CASCADE,
    action STRING(50) NOT NULL,
    user_id INT8 REFERENCES users (id) ON DELETE CASCADE,
    writed_at INT8 DEFAULT (extract(epoch FROM current_timestamp()) * 1000)::INT8
);

CREATE INDEX idx_activity_field_user ON activity (activity_field_id, user_id);