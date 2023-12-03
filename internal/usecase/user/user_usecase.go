package user

import (
	"context"

	authdto "github.com/vnnyx/betty-BE/internal/delivery/http/dto/auth"
	userdto "github.com/vnnyx/betty-BE/internal/delivery/http/dto/user"
)

type UserUC interface {
	CreateOwner(ctx context.Context, req *userdto.CreateOwnerRequest) (*authdto.AuthResponse, error)
}
