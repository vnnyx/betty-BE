CREATE TABLE activity_detail (
    id INT8 PRIMARY KEY DEFAULT unique_rowid(),
    changed_id INT8 DEFAULT NULL,
    activity_id INT8 REFERENCES activity (id) ON DELETE CASCADE,
    old_value STRING DEFAULT NULL,
    new_value STRING DEFAULT NULL
);

CREATE INDEX idx_activity_changed ON activity_detail (changed_id, activity_id);