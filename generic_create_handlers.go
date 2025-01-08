package gocrud

import (
	"context"

	"github.com/dargasht/gocrud/internal/model"
	"github.com/gofiber/fiber/v2"
)

func NewCreateAdminJSONHandler[T CReq[U], U CRepo, R CRes](
	c *fiber.Ctx,
	h *HandlerConfig,
	source string,
	createFunc func(context.Context, U) (R, error),
) error {

	if err := EnsureAdmin(source, h, c); err != nil {
		return err
	}

	var req T
	if err := c.BodyParser(&req); err != nil {
		return NewJSONError(source+" JSON Parse Error", err)
	}

	if err := model.Validate.Struct(&req); err != nil {
		return NewValidationError(source+" Validation Error", err)
	}

	res, err := createFunc(c.Context(), req.ToRepo())
	if err != nil {
		return NewCreateError(source+" Database Error", err)
	}

	return c.Status(fiber.StatusCreated).JSON(NewRes(res, Success, fiber.StatusCreated))

}
