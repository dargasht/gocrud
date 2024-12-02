package gocrud

import (
	"context"

	"github.com/dargasht/gocrud/helper"
	"github.com/gofiber/fiber/v2"
)

func NewDeleteAdminJSONHandler(
	c *fiber.Ctx,
	h *HandlerConfig,
	source string,
	updateFunc func(context.Context, int32) (int64, error),
) error {

	if err := EnsureAdmin(source, h, c); err != nil {
		return err
	}

	id, _ := c.ParamsInt("id")

	rowsAffected, err := updateFunc(c.Context(), int32(id))
	if err != nil {
		return NewUpdateError(source+" Update Error", err.Error())
	}

	if rowsAffected == 0 {
		return NewNotFoundError(source+" Not Found", "not found")
	}

	return c.Status(fiber.StatusCreated).JSON(NewRes([]int{}, helper.Success, fiber.StatusCreated))
}
