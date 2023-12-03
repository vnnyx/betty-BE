package entity

type UserGroup struct {
	ID         int64 `json:"id"`
	UserID     int64 `json:"-"`
	User       *User `json:"user,omitempty"`
	UserRoleID int64 `json:"-"`
	Role       *Role `json:"role,omitempty"`
}
