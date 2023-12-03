package auth

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/vnnyx/betty-BE/internal/delivery/http/dto"
	authdto "github.com/vnnyx/betty-BE/internal/delivery/http/dto/auth"
	"github.com/vnnyx/betty-BE/internal/enums"
	"github.com/vnnyx/betty-BE/internal/pkg/utils/authutils"
	"github.com/vnnyx/betty-BE/internal/usecase/auth"
	"golang.org/x/oauth2"
	"google.golang.org/api/option"
	"google.golang.org/api/people/v1"
)

type AuthHandler struct {
	authUsecase auth.AuthUC
	googleAuth  *oauth2.Config
}

func NewAuthHandler(authUsecase auth.AuthUC, googleAuth *oauth2.Config) *AuthHandler {
	return &AuthHandler{
		authUsecase: authUsecase,
		googleAuth:  googleAuth,
	}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req authdto.LoginRequest
	err := c.BodyParser(&req)
	if err != nil {
		return err
	}

	res, err := h.authUsecase.Login(c.Context(), &req)
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

func (h *AuthHandler) RefreshToken(c *fiber.Ctx) error {
	var req authdto.RefreshTokenRequest
	err := c.BodyParser(&req)
	if err != nil {
		return err
	}

	res, err := h.authUsecase.GetNewAccessToken(c.Context(), &req)
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

func (h *AuthHandler) Logout(c *fiber.Ctx) error {
	return nil
}

func (h *AuthHandler) GoogleSign(c *fiber.Ctx) error {
	state, err := authutils.GenerateRandomString(32)
	if err != nil {
		return err
	}
	url := h.googleAuth.AuthCodeURL(state, oauth2.AccessTypeOffline)
	return c.Redirect(url)
}

func (h *AuthHandler) GoogleCallback(c *fiber.Ctx) error {
	code := c.Query("code")
	token, err := h.googleAuth.Exchange(c.Context(), code)
	if err != nil {
		return err
	}

	client := h.googleAuth.Client(context.Background(), token)
	service, err := people.NewService(c.Context(), option.WithHTTPClient(client))
	if err != nil {
		return err
	}

	peopleResource, err := service.People.Get("people/me").PersonFields("emailAddresses,names,addresses,phoneNumbers,photos").Do()
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusOK).JSON(dto.WebResponse{
		Code:   fiber.StatusOK,
		Status: enums.StatusOK,
		Data: map[string]any{
			"people": peopleResource,
			"token":  token,
		},
		Error: nil,
	})
}
