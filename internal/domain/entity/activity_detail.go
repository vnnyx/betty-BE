package entity

type ActivityDetail struct {
	ID         int64     `json:"id"`
	ActivityID int64     `json:"activity_id"`
	ChangedID  *int64    `json:"changed_id"`
	Activity   *Activity `json:"activity,omitempty"`
	OldValue   *string   `json:"old_value"`
	NewValue   *string   `json:"new_value"`
}
