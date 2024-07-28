package product

import (
	"net/http"
	"strconv"

	"github.com/gofiber/fiber/v2"
	"github.com/ohmtanawin02/go-postgres-basic/have-api/models"
)

func GetProductsHandler(c *fiber.Ctx) error {
	name := c.Query("name")
	page, err := strconv.Atoi(c.Query("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}
	limit, err := strconv.Atoi(c.Query("limit", "10"))
	if err != nil || limit < 1 {
		limit = 10
	}

	productsResponse, err := GetProducts(c.Context(), name, page, limit)
	if err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(productsResponse)
}

func GetProductHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid product ID"})
	}

	product, err := GetProduct(c.Context(), id)
	if err != nil {
		return c.Status(http.StatusNotFound).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(product)
}

func CreateProductHandler(c *fiber.Ctx) error {
	product := new(models.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	if err := CreateProduct(c.Context(), product); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(http.StatusCreated).JSON(product)
}

func UpdateProductHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid product ID"})
	}

	product := new(models.Product)
	if err := c.BodyParser(product); err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "cannot parse JSON"})
	}

	if err := UpdateProduct(c.Context(), id, product); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(product)
}

func DeleteProductHandler(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "invalid product ID"})
	}

	if err := DeleteProduct(c.Context(), id); err != nil {
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.SendStatus(http.StatusNoContent)
}
