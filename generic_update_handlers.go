package gocrud

import (
	"context"

	"github.com/dargasht/gocrud/helper"
	"github.com/dargasht/gocrud/model"
	"github.com/gofiber/fiber/v2"
)

func NewUpdateAdminJSONHandler[T UReq[U], U URepo](
	c *fiber.Ctx,
	h *HandlerConfig,
	source string,
	updateFunc func(context.Context, U) (int64, error),
) error {

	if err := EnsureAdmin(source, h, c); err != nil {
		return err
	}

	var req T
	if err := c.BodyParser(&req); err != nil {
		return NewJsonError(source+" JSON Parse Error", err.Error())
	}
	if err := model.Validate.Struct(&req); err != nil {
		return NewValidationError(source+" Validation Error", err.Error())
	}

	id, _ := c.ParamsInt("id")
	a, ok := req.SetID(int32(id)).(T)
	if !ok {
		return NewNotFoundError(source+" Not Found", "not found")
	}

	rowsAffected, err := updateFunc(c.Context(), a.ToRepo())
	if err != nil {
		return NewUpdateError(source+" Update Error", err.Error())
	}

	if rowsAffected == 0 {
		return NewNotFoundError(source+" Not Found", "not found")
	}

	return c.Status(fiber.StatusCreated).JSON(NewRes([]int{}, helper.Success, fiber.StatusCreated))
}
