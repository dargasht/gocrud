package gocrud

import (
	"context"

	"github.com/dargasht/gocrud/helper"
	"github.com/dargasht/gocrud/model"
	"github.com/gofiber/fiber/v2"
)

type CreateFunc func(context.Context, CRepo) (CRes, error)

func NewCreateAdminJSONHandler[T CReq[U], U CRepo](c *fiber.Ctx, h *HandlerConfig, source string, createFunc CreateFunc) error {

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

	res, err := createFunc(c.Context(), req.ToRepo())
	if err != nil {
		return NewCreateErrro(source+" Database Error", err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(NewRes(res, helper.Success, fiber.StatusCreated))

}
