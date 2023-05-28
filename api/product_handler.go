package api

import (
	"errors"
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

	if err := h.store.Insert(c.Context(), product); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to insert product",
		})
	}

	return c.JSON(fiber.StatusOK, product)
}

func (h *ProductHandler) HandleGetProducts(c *fiber.Ctx) error {
	products, err := h.store.GetAll(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to retreive products",
		})
	}
	return c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) HandleGetProductByID(c *fiber.Ctx) error {
	id := c.Params("id")

	product, err := h.store.GetByID(c.Context, id)
	if err != nil {
		if errors.Is(err, ErrProductNotFound) {
			return c.Status(http.StatusNotFound).JSON(fiber.Map{
				"message": "Product not found",
			})
		}
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{
			"message": "Internal server error"})
	}

	return c.JSON(product)
}
