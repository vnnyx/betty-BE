package entity

type RoleScope struct {
	ID      int64  `json:"id"`
	RoleID  int64  `json:"-"`
	Role    *Role  `json:"role,omitempty"`
	ScopeID int64  `json:"-"`
	Scope   *Scope `json:"scope,omitempty"`
}
