package router

import (
	"github.com/dargasht/gocrud"
	"github.com/dargasht/gocrud/internal/middleware"
	"github.com/gofiber/fiber/v2"
)

type routeConfig struct {
	r             fiber.Router
	handlerConfig gocrud.HandlerConfig
	jwt           middleware.Middleware
}
