package handlers

import (
	"errors"
	"net/http"

	"github.com/e-politica/api/pkg/session"
	"github.com/e-politica/api/routes"
	"github.com/e-politica/api/routes/v1/user/repository"
	"github.com/gofiber/fiber/v2"
)

func PostLoginGoogle(tools routes.Tools) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var credentials string
		if err := c.BodyParser(&credentials); err != nil {
			tools.Logger.Error.Println(err)
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "could not parse request body"})
		}

		userSession, err := repository.LoginGoogle(c.Context(), tools.Db, credentials)
		if err != nil {
			code := http.StatusBadRequest
			if err != session.ErrSessionNotFound &&
				err != repository.ErrInvalidJwt {
				tools.Logger.Error.Println(err)
				err = errors.New("internal server error")
				code = http.StatusBadRequest
			}
			return c.Status(code).JSON(fiber.Map{"error": err.Error()})
		}

		return c.Status(http.StatusOK).JSON(userSession)
	}
}
