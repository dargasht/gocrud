package gocrud

import (
	"github.com/dargasht/gocrud/database/repo"
	"github.com/dargasht/gocrud/service"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
)

type HandlerConfig struct {
	DB     *repo.Queries
	Logger *zap.Logger
}

func NewHandlerConfig(db *repo.Queries, logger *zap.Logger) *HandlerConfig {
	return &HandlerConfig{
		DB:     db,
		Logger: logger,
	}
}

// ------------------------------------------------------------------
// ------------------------Response----------------------------------
// ------------------------------------------------------------------

type Res[T CRes] struct {
	Data    T      `json:"data"`
	Message string `json:"message"`
	Status  int    `json:"status"`
}

func NewRes[T CRes](data T, message string, status int) Res[T] {
	return Res[T]{
		Data:    data,
		Message: message,
		Status:  status,
	}
}

type ResWithMeta struct {
	Data    any    `json:"data"`
	Message string `json:"message"`
	Status  int    `json:"status"`
	Meta    Meta   `json:"meta,omitempty"`
}

func NewResWithMeta(data any, message string, status int, meta Meta) ResWithMeta {
	return ResWithMeta{
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
	if user.Role == "admin" || user.Role == "owner" {
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
