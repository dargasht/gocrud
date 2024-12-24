package gocrud

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgconn"
)

// Some Simple errors
var ErrNotFound error = errors.New("no data found")

// -------------------------------------------------------------------
// --------------------	Our custom error -----------------------------
// -------------------------------------------------------------------

// This is our custom error
// It is a struct that contains the error code, message, source, internal message and args
// Code is code (e.g. 400), Message is what we show to the user,
// Source is where the error came from, InternalMessage is what we use for logging,
// Args is any argument you want to add to InternalMessage or Message
// (you should add to whatever you want in the function that creates the error, don't fuck up plz)
// At this moment, all the error messages that are send to the user are in Persian,
// Maybe in the future we will add more languages but for now we only have Persian
//
// Regarding source, you should have some sources constants in your code like this
// const (
//  userhandlerDb = "user handler -> db"
//  userhandlerApi = "user handler -> api"
// )

type CRUDError struct {
	Code            int      `json:"code"`
	Message         string   `json:"message"`
	Source          string   `json:"-"`
	InternalMessage string   `json:"-"`
	Args            []string `json:"-"`
}

// This is what make this struct error
func (e *CRUDError) Error() string {
	return e.Message
}

// Only use when there is no error and want to create a new error
// I barely use this one
func NewCRUDError(code int, message string, source string, internalMessage string, args ...string) *CRUDError {
	return &CRUDError{
		Code:            code,
		Message:         message,
		Source:          source,
		InternalMessage: internalMessage,
		Args:            args,
	}
}

// Used for 404
func NewNotFoundError(source string, err error, args ...string) *CRUDError {
	return &CRUDError{
		Code:            fiber.StatusNotFound,
		Message:         NoDataFound,
		Source:          source,
		InternalMessage: err.Error(),
		Args:            args,
	}
}

// ------------------------------------------------------------------
// ------------------- Database Errors ------------------------------
// ------------------------------------------------------------------
func NewDBError(source string, err error, args ...string) *CRUDError {
	return &CRUDError{
		Code:            fiber.StatusInternalServerError,
		Message:         InternalServerError,
		Source:          source,
		InternalMessage: err.Error(),
		Args:            args,
	}
}

func NewCreateError(source string, err error, args ...string) *CRUDError {

	var pgErr *pgconn.PgError

	switch {
	case errors.As(err, &pgErr):
		switch err.(*pgconn.PgError).Code {
		case "23505":
			return &CRUDError{
				Code:            fiber.StatusBadRequest,
				Message:         DataNotUnique,
				Source:          source,
				InternalMessage: err.Error(),
				Args:            args,
			}

		default:
			return &CRUDError{
				Code:            fiber.StatusBadRequest,
				Message:         CreateError,
				Source:          source,
				InternalMessage: err.Error(),
				Args:            args,
			}
		}

	default:
		return &CRUDError{
			Code:            fiber.StatusBadRequest,
			Message:         CreateError,
			Source:          source,
			InternalMessage: err.Error(),
			Args:            args,
		}
	}
}

func NewUpdateError(source string, err error, args ...string) *CRUDError {
	return &CRUDError{
		Code:            fiber.StatusInternalServerError,
		Message:         UpdateError,
		Source:          source,
		InternalMessage: err.Error(),
		Args:            args,
	}
}

func NewDeleteError(source string, err error, args ...string) *CRUDError {
	return &CRUDError{
		Code:            fiber.StatusInternalServerError,
		Message:         DeleteError,
		Source:          source,
		InternalMessage: err.Error(),
		Args:            args,
	}
}

// ------------------------------------------------------------------
// ---------------- JSON and Validation Errors ----------------------
// ------------------------------------------------------------------
func NewJSONError(source string, err error, args ...string) *CRUDError {
	return &CRUDError{
		Code:            fiber.StatusBadRequest,
		Message:         InvalidJsonBody,
		Source:          source,
		InternalMessage: err.Error(),
		Args:            args,
	}
}

func NewValidationError(source string, err error, args ...string) *CRUDError {
	return &CRUDError{
		Code:            fiber.StatusBadRequest,
		Message:         InvalidRequest + err.Error(),
		Source:          source,
		InternalMessage: err.Error(),
		Args:            args,
	}
}

// ------------------------------------------------------------------
// ------------------- Other Errors ---------------------------------
// ------------------------------------------------------------------
func NewPermissionError(source string, internalMessage string, args ...string) *CRUDError {
	return &CRUDError{
		Code:            fiber.StatusForbidden,
		Message:         PermissionNotAllowed,
		Source:          source,
		InternalMessage: internalMessage,
		Args:            args,
	}
}

func NewInvalidTokenError(source string, internalMessage string) *CRUDError {
	return &CRUDError{
		Code:            fiber.StatusUnauthorized,
		Message:         InvalidToken,
		Source:          source,
		InternalMessage: internalMessage,
	}
}
