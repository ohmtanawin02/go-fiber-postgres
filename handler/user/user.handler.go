package user

// import (
// 	"database/sql"

// 	"github.com/gofiber/fiber/v2"
// 	"github.com/ohmtanawin02/go-postgres-basic/have-api/models"
// 	"golang.org/x/crypto/bcrypt"
// )

// func Register(c *fiber.Ctx) error {
// 	var user models.User
// 	if err := c.BodyParser(&user); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
// 	}

// 	err := CreateUser(&user)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error creating user"})
// 	}

// 	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User created"})
// }

// func Login(c *fiber.Ctx) error {
// 	var reqUser models.User
// 	if err := c.BodyParser(&reqUser); err != nil {
// 		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
// 	}

// 	user, err := GetUserByUsername(reqUser.Username)
// 	if err != nil {
// 		if err == sql.ErrNoRows {
// 			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
// 		}
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error querying database"})
// 	}

// 	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(reqUser.Password))
// 	if err != nil {
// 		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
// 	}

// 	token, err := middlewares.GenerateJWT(user.ID)
// 	if err != nil {
// 		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Error generating JWT"})
// 	}

// 	return c.JSON(fiber.Map{"token": token})
// }
