package entity

type Activity struct {
	ID              int64  `json:"id"`
	ActivityFieldID int64  `json:"activity_field_id"`
	Action          string `json:"action"`
	UserID          int64  `json:"-"`
	User            *User  `json:"user,omitempty"`
	WritedAt        int64  `json:"writed_at"`
}
