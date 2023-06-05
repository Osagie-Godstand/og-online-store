package api

import (
	"net/http"

	"github.com/Osagie-Godstand/og-online-store/db"
	"github.com/Osagie-Godstand/og-online-store/types"
	"github.com/gofiber/fiber/v2"
)

type ProductHandler struct {
	store db.ProductStorer
}

func NewProductHandler(pStore db.ProductStorer) *ProductHandler {
	return &ProductHandler{
		store: pStore,
	}
}

func (h *ProductHandler) HandlePostProduct(c *fiber.Ctx) error {
	productReq := &types.CreateProductRequest{}
	if err := c.BodyParser(productReq); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid request payload",
		})
	}

	product, err := types.NewProductFromRequest(productReq)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to create product",
		})
	}
	insertedProduct, err := h.store.InsertProduct(c.Context(), product)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to insert product",
		})
	}

	return c.JSON(insertedProduct)

}

func (h *ProductHandler) HandleGetProductByID(c *fiber.Ctx) error {
	id := c.Params("id")
	product, err := h.store.GetByID(c.Context(), id)
	if err != nil {
		return fiber.NewError(http.StatusInternalServerError, "Internal server error")
	}

	return c.JSON(product)
}
