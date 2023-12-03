package entity

type User struct {
	ID           int64           `json:"id"`
	IsSuperAdmin bool            `json:"is_super_admin"`
	IsAdmin      bool            `json:"is_admin"`
	Email        string          `json:"email"`
	Password     string          `json:"-"`
	FullName     string          `json:"full_name"`
	PhoneNumber  string          `json:"phone_number"`
	CompanyID    int64           `json:"-"`
	Company      *Company        `json:"company,omitempty"`
	FranchiseID  int64           `json:"-"`
	Franchise    *Franchise      `json:"franchise,omitempty"`
	IsActive     bool            `json:"is_active"`
	PhotoID      int64           `json:"-"`
	Photo        *AttachmentFile `json:"photo,omitempty"`
	RefreshToken string          `json:"refresh_token"`
	SharedSecret string          `json:"shared_secret"`
}
