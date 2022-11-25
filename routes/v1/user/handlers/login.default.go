package handlers

import (
	"errors"
	"net/http"

	"github.com/e-politica/api/models/v1/user"
	"github.com/e-politica/api/pkg/session"
	"github.com/e-politica/api/routes"
	"github.com/e-politica/api/routes/v1/user/repository"
	"github.com/gofiber/fiber/v2"
)

func PostLoginDefault(tools routes.Tools) fiber.Handler {
	return func(c *fiber.Ctx) error {
		var params user.LoginDefaultParams
		if err := c.BodyParser(&params); err != nil {
			tools.Logger.Error.Println(err)
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "could not parse request body."})
		}

		if err := params.Validate(); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		userSession, err := repository.LoginDefault(c.Context(), tools.Db, params)
		if err != nil {
			code := http.StatusBadRequest
			if err != session.ErrSessionNotFound &&
				err != repository.ErrInexistentAccount &&
				err != repository.ErrPasswordsDontMatch {
				tools.Logger.Error.Println(err)
				err = errors.New("internal server error")
				code = http.StatusInternalServerError
			}
			return c.Status(code).JSON(fiber.Map{"error": err.Error()})
		}

		return c.Status(http.StatusOK).JSON(userSession)
	}
}
