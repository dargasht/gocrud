package gocrud

import (
	"github.com/dargasht/gocrud/internal/database/repo"
	"github.com/dargasht/gocrud/internal/service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

// Example of what a handler config should look like
// It is not recommanded to use this just write your own
type HandlerConfig struct {
	DB     *repo.Queries
	Logger *zap.Logger
}

// Creates a new handler config
// This is not recommanded to use this just write your own
func NewHandlerConfig(db *repo.Queries, logger *zap.Logger) *HandlerConfig {
	return &HandlerConfig{
		DB:     db,
		Logger: logger,
	}
}

// ------------------------------------------------------------------
// ------------------------Response----------------------------------
// ------------------------------------------------------------------

// Standard response suitable for most handlers
type Res[T any] struct {
	Data    T      `json:"data"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func NewRes[T any](data T, message string, status int) Res[T] {
	return Res[T]{
		Data:    data,
		Message: message,
		Status:  status,
	}
}

// Standard response suitable for handlers who do pagination
// In go 1.24 when the new omitzero json tag will be introduced
// we can merge the 2 responses probably
type ResWithMeta[T any] struct {
	Data    T      `json:"data"`
	Message string `json:"message"`
	Status  int    `json:"status"`
	Meta    Meta   `json:"meta,omitempty"`
}

func NewResWithMeta[T any](data T, message string, status int, meta Meta) ResWithMeta[T] {
	return ResWithMeta[T]{
		Data:    data,
		Message: message,
		Status:  status,
		Meta:    meta,
	}
}

type Meta struct {
	CurrentPage int `json:"current_page"`
	LastPage    int `json:"last_page"`
	Total       int `json:"total"`
	FromIndex   int `json:"from_index"`
	ToIndex     int `json:"to_index"`
	PerPage     int `json:"per_page"`
}

// ------------------------------------------------------------------
// ------------------------Pagination--------------------------------
// ------------------------------------------------------------------

// GetPagination returns page, limit and offset
func GetPagination(c *fiber.Ctx) (page int, limit int, offset int) {

	page = c.QueryInt("page", 1)
	if page < 1 {
		page = 1
	}

	limit = c.QueryInt("limit", 30)
	if limit < 1 || limit > 30 {
		limit = 30
	}

	offset = (page - 1) * limit

	return page, limit, offset
}

// GetPaginationNoLimit returns page, limit and offset
// Use in rare cases that user wants to get alot of data
func GetPaginationNoLimit(c *fiber.Ctx) (page int, limit int, offset int) {

	page = c.QueryInt("page", 1)
	if page < 1 {
		page = 1
	}

	limit = c.QueryInt("limit", 30)
	if limit < 1 {
		limit = 30
	}

	offset = (page - 1) * limit

	return page, limit, offset

}

// GetPaginationMeta returns pagination meta
// For putting in the response
func GetPaginationMeta(page int, limit int, size int) Meta {

	last_page := size / limit

	if size < limit {
		last_page++
	}

	return Meta{
		CurrentPage: page,
		LastPage:    last_page,
		Total:       size,
		FromIndex:   (page-1)*limit + 1,
		ToIndex:     (page-1)*limit + limit,
		PerPage:     limit,
	}

}

// ------------------------------------------------------------------
// ------------------------Auth helpers------------------------------
// ------------------------------------------------------------------

// EnsureAdmin checks if the user is admin
// You can make this in a middleware
// This is just an example for a system who has admin role
// You can use it as a template for your system and roles
func EnsureAdmin(source string, h *HandlerConfig, c *fiber.Ctx) error {

	user, err := Authenticate(source, h, c)
	if err != nil {
		return err
	}

	if !userHaveAdminAccess(user) {
		return NewPermissionError(source, "permission denied")
	}

	return nil
}

// This is Very usefull for authenticating the user
// Usefull for most crud applications
func Authenticate(source string, h *HandlerConfig, c *fiber.Ctx) (repo.User, error) {

	user := repo.User{}
	t, ok := c.Locals("user").(*jwt.Token)
	if !ok {
		return user, NewInvalidTokenError(source, "invalid token in here t, ok := c.Locals(\"user\").(*jwt.Token)")
	}

	claims, ok := t.Claims.(jwt.MapClaims)
	if !ok {
		return user, NewInvalidTokenError(source, "invalid token in here claims, ok := t.Claims.(jwt.MapClaims)")
	}

	u, ok := claims["user_id"].(float64)
	if !ok {
		return user, NewInvalidTokenError(source, "invalid token in here u, ok := claims[user_id].(float64)")
	}

	userID := int32(u)
	user, err := h.DB.GetUserByID(c.Context(), userID)
	if err != nil {
		return user, NewInvalidTokenError(source, err.Error())
	}

	return user, nil
}

// For cases that a request can contain or not contain a token
// Use this to get the user id
func GetUserIDFromJWT(c *fiber.Ctx) int32 {

	tokenString, err := service.GetJWTFromHeader(c, "Bearer")
	if err != nil {
		return 0
	}

	claims, err := service.DecodeJWT(tokenString)
	if err != nil {
		return 0
	}

	u, ok := claims["user_id"].(float64)
	if !ok {
		return 0
	}

	return int32(u)
}

// ------------------------------------------------------------------
// ------------------------Check Roles-------------------------------
// ------------------------------------------------------------------
func userHaveAdminAccess(user repo.User) bool {
	if user.Role == "admin" {
		return true
	}
	return false
}

func selfOrAdminAccess(user repo.User, id int32) bool {
	if user.ID == id || userHaveAdminAccess(user) {
		return true
	}
	return false
}
