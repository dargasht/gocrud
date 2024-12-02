package middleware

import "github.com/gofiber/fiber/v2"

type Middleware func(*fiber.Ctx) error
