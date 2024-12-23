package gocrud

import (
	"errors"
	"strconv"

	"github.com/dargasht/gocrud/internal/helper"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

// MY custom error
type CRUDError struct {
	Code            int    `json:"code"`
	Message         string `json:"message"`
	Source          string `json:"-"`
	InternalMessage string `json:"-"`
}

func (e *CRUDError) Error() string {
	return e.Message
}

func NewCRUDError(code int, message string, source string, internalMessage string) *CRUDError {
	return &CRUDError{
		Code:            code,
		Message:         message,
		Source:          source,
		InternalMessage: internalMessage,
	}
}

func CustomErrorHandler(log *zap.Logger) fiber.ErrorHandler {
	return func(c *fiber.Ctx, err error) error {

		var myErr *CRUDError

		if errors.As(err, &myErr) {
			log.Info("error in: "+myErr.Source, zap.String("error: ", myErr.InternalMessage))

			return c.Status(myErr.Code).JSON(&CRUDError{
				Code:    myErr.Code,
				Message: myErr.Message,
			})
		} else if e, ok := err.(*fiber.Error); ok {
			log.Warn("fiber error " + strconv.Itoa(e.Code) + e.Message)
			return c.Status(fiber.StatusNotFound).JSON(&CRUDError{
				Code:    fiber.StatusNotFound,
				Message: "Not Found",
			})

		}

		log.Error("Unknown Error", zap.String("error", err.Error()))
		// fmt.Println("Unknown Error: " + err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(&CRUDError{
			Code:    fiber.StatusInternalServerError,
			Message: "Internal Server Error",
		})
	}
}

// -------------------------------------------------------------------
// -----------------------Custom Errors-------------------------------
// -------------------------------------------------------------------
func NewInvalidTokenError(source string, internalMessage string) *CRUDError {
	return &CRUDError{
		Code:            fiber.StatusUnauthorized,
		Message:         helper.InvalidToken,
		Source:          source,
		InternalMessage: internalMessage,
	}
}

func NewPermissionError(source string, internalMessage string) *CRUDError {
	return &CRUDError{
		Code:            fiber.StatusForbidden,
		Message:         helper.PermissionNotAllowed,
		Source:          source,
		InternalMessage: internalMessage,
	}
}

func NewJsonError(source string, internalMessage string) *CRUDError {
	return &CRUDError{
		Code:            fiber.StatusUnprocessableEntity,
		Message:         helper.InvalidJsonBody,
		Source:          source,
		InternalMessage: internalMessage,
	}
}

func NewValidationError(source string, internalMessage string) *CRUDError {
	return &CRUDError{
		Code:            fiber.StatusUnprocessableEntity,
		Message:         helper.InvalidRequest,
		Source:          source,
		InternalMessage: internalMessage,
	}
}

func NewCreateErrro(source string, internalMessage string) *CRUDError {
	return &CRUDError{
		Code:            fiber.StatusInternalServerError,
		Message:         helper.CreateError,
		Source:          source,
		InternalMessage: internalMessage,
	}
}

func NewUpdateError(source string, internalMessage string) *CRUDError {
	return &CRUDError{
		Code:            fiber.StatusInternalServerError,
		Message:         helper.UpdateError,
		Source:          source,
		InternalMessage: internalMessage,
	}
}

func NewNotFoundError(source string, internalMessage string) *CRUDError {
	return &CRUDError{
		Code:            fiber.StatusNotFound,
		Message:         helper.NotFound,
		Source:          source,
		InternalMessage: internalMessage,
	}
}
