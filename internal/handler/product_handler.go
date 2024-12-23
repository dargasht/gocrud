package handler

import (
	"github.com/dargasht/gocrud"
	"github.com/dargasht/gocrud/internal/database/repo"
	"github.com/dargasht/gocrud/internal/model"
	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	gocrud.HandlerConfig
}

func NewProductHandler(config gocrud.HandlerConfig) *ProductHandler {
	return &ProductHandler{config}
}

//-------------------------------------------------------------------
//---------------------------Handlers--------------------------------
//-------------------------------------------------------------------

func (h *ProductHandler) Create(c *fiber.Ctx) error {

	return gocrud.NewCreateAdminJSONHandler[model.ProductCReq, repo.CreateProductParams, repo.Product](
		c,
		&h.HandlerConfig,
		"Product",
		h.DB.CreateProduct,
	)
}

func (h *ProductHandler) Update(c *fiber.Ctx) error {

	return gocrud.NewUpdateAdminJSONHandler[model.ProductUReq, repo.UpdateProductParams](
		c,
		&h.HandlerConfig,
		"Product",
		h.DB.UpdateProduct,
	)
}

func (h *ProductHandler) Delete(c *fiber.Ctx) error {

	return gocrud.NewDeleteAdminJSONHandler(c, &h.HandlerConfig, "Product", h.DB.DeleteProduct)
}
