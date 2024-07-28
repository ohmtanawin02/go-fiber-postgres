package product

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterRoutes(app *fiber.App) {
	route := app.Group("/products")

	route.Get("/", GetProductsHandler)
	route.Get("/:id", GetProductHandler)
	route.Post("/", CreateProductHandler)
	route.Put("/:id", UpdateProductHandler)
	route.Delete("/:id", DeleteProductHandler)
}
