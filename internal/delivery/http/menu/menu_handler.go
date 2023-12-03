package menu

import (
	"github.com/gofiber/fiber/v2"
	"github.com/vnnyx/betty-BE/internal/delivery/http/dto"
	menudto "github.com/vnnyx/betty-BE/internal/delivery/http/dto/menu"
	"github.com/vnnyx/betty-BE/internal/enums"
	"github.com/vnnyx/betty-BE/internal/usecase/menu"
)

type MenuHandler struct {
	menuUC menu.MenuUC
}

func NewMenuHandler(menuUC menu.MenuUC) *MenuHandler {
	return &MenuHandler{
		menuUC: menuUC,
	}
}

func (h *MenuHandler) AddMenu(c *fiber.Ctx) error {
	var req menudto.CreateMenuRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	res, err := h.menuUC.CreateMenu(c.Context(), &req)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(dto.WebResponse{
		Code:   fiber.StatusCreated,
		Status: enums.StatusCreated,
		Data:   res,
		Error:  nil,
	})
}

func (h *MenuHandler) AddBaseProduct(c *fiber.Ctx) error {
	var req menudto.CreateBaseProductRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	res, err := h.menuUC.CreateBaseProduct(c.Context(), &req)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(dto.WebResponse{
		Code:   fiber.StatusCreated,
		Status: enums.StatusCreated,
		Data:   res,
		Error:  nil,
	})
}

func (h *MenuHandler) AddMenuRecipe(c *fiber.Ctx) error {
	var req menudto.CreateMenuRecipeRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	res, err := h.menuUC.CreateMenuRecipe(c.Context(), &req)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(dto.WebResponse{
		Code:   fiber.StatusCreated,
		Status: enums.StatusCreated,
		Data:   res,
		Error:  nil,
	})
}

func (h *MenuHandler) AddMenuCategory(c *fiber.Ctx) error {
	var req menudto.CreateMenuCategoryRequest
	if err := c.BodyParser(&req); err != nil {
		return err
	}

	res, err := h.menuUC.CreateMenuCategory(c.Context(), &req)
	if err != nil {
		return err
	}

	return c.Status(fiber.StatusCreated).JSON(dto.WebResponse{
		Code:   fiber.StatusCreated,
		Status: enums.StatusCreated,
		Data:   res,
	})
}
