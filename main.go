package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/ohmtanawin02/go-postgres-basic/config"
	"github.com/ohmtanawin02/go-postgres-basic/handler/product"
)

func main() {
	app := fiber.New()
	config.InitDB()
	product.RegisterRoutes(app)
	// user.RegisterRoutes(app)

	log.Println("Server is running on port 8080")
	log.Fatal(app.Listen(":8080"))
}
