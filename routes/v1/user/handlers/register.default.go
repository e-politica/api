package handlers

import (
	"log"
	"net/http"

	"github.com/e-politica/api/models/v1/user"
	"github.com/e-politica/api/pkg/database"
	"github.com/e-politica/api/routes/v1/user/repository"
	"github.com/gofiber/fiber/v2"
)

func PostRegisterDefault(db *database.Db) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var body user.RegisterDefault
		if err := c.BodyParser(&body); err != nil {
			log.Println(err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "could not parse request body"})
		}

		if err := body.Validate(); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		_, err := repository.RegisterDefault(c.Context(), db, body)
		if err != nil {
			log.Println(err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.SendStatus(http.StatusOK)
	}
}
