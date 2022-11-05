package handlers

import (
	"net/http"

	"github.com/e-politica/api/models/v1/user"
	"github.com/e-politica/api/routes"
	"github.com/e-politica/api/routes/v1/user/repository"
	"github.com/gofiber/fiber/v2"
)

func PostRegisterDefault(tools routes.Tools) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var params user.RegisterDefault
		if err := c.BodyParser(&params); err != nil {
			tools.Logger.Error.Println(err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "could not parse request body"})
		}

		if err := params.Validate(); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		session, err := repository.RegisterDefault(c.Context(), tools.Db, params)
		if err != nil {
			if err != repository.ErrExistentAccount {
				tools.Logger.Error.Println(err)
			}
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.Status(http.StatusCreated).JSON(session)
	}
}
