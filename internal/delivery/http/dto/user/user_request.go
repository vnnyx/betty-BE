package user

import "github.com/vnnyx/betty-BE/internal/delivery/http/dto/company"

type CreateOwnerRequest struct {
	Fullname             string                        `json:"fullname" validate:"required"`
	PhoneNumber          string                        `json:"phone_number" validate:"required,number"`
	Email                string                        `json:"email" validate:"required,email"`
	Password             string                        `json:"password" validate:"required,min=8,max=32"`
	PasswordConfirmation string                        `json:"password_confirmation" validate:"required,min=8,max=32,eqfield=Password"`
	Company              *company.CreateCompanyRequest `json:"company" validate:"required"`
	Photo                *string                       `json:"photo,omitempty"`
}
