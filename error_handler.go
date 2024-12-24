package gocrud

import (
	"errors"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// This function for handling errors that are returned from the handlers
// Just use it
func CRUDErrorHandler(log *zap.Logger) fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {

		var crudErr *CRUDError
		e, ok := err.(*fiber.Error)

		switch {
		case errors.As(err, &crudErr):

			log.Info(
				"error in: "+crudErr.Source,
				zap.String("error: ", crudErr.InternalMessage),
				zap.String("args: ", strings.Join(crudErr.Args, " ")),
			)
			return c.Status(crudErr.Code).JSON(&CRUDError{
				Code:    crudErr.Code,
				Message: crudErr.Message,
			})

		case ok:
			log.Warn("fiber error " + strconv.Itoa(e.Code) + e.Message)
			return c.Status(fiber.StatusNotFound).JSON(&CRUDError{
				Code:    fiber.StatusNotFound,
				Message: "Not Found",
			})
		default:
			log.Error("Unknown Error", zap.String("error", err.Error()))
		}

		return c.Status(fiber.StatusInternalServerError).JSON(&CRUDError{
			Code:    fiber.StatusInternalServerError,
			Message: "Internal Server Error",
		})
	}
}
