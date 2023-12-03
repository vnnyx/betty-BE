package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vnnyx/betty-BE/internal/delivery/http/dto"
	userdto "github.com/vnnyx/betty-BE/internal/delivery/http/dto/user"
	"github.com/vnnyx/betty-BE/internal/enums"
	"github.com/vnnyx/betty-BE/internal/usecase/user"
)

type UserHandler struct {
	userUC user.UserUC
}

func NewUserHandler(userUC user.UserUC) *UserHandler {
	return &UserHandler{
		userUC: userUC,
	}
}

func (h *UserHandler) CreateOwner(c *fiber.Ctx) error {
	var req userdto.CreateOwnerRequest
	err := c.BodyParser(&req)
	if err != nil {
		return err
	}

	res, err := h.userUC.CreateOwner(c.Context(), &req)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.WebResponse{
		Code:   fiber.StatusOK,
		Status: enums.StatusOK,
		Data:   res,
		Error:  nil,
	})
}
