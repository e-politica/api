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

func PutChangeInfo(tools routes.Tools) fiber.Handler {
	return func(c *fiber.Ctx) error {
		access := c.Get("Authorization")
		if access == "" {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "must provide 'Authorization' header"})
		}

		var params user.ChangeInfoParams
		if err := c.BodyParser(&params); err != nil {
			tools.Logger.Error.Println(err)
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "could not parse request body"})
		}

		if err := params.Validate(); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		err := repository.ChangeInfo(c.Context(), tools.Db, access, params)
		if err != nil {
			code := http.StatusBadRequest
			if err != session.ErrSessionNotFound &&
				err != repository.ErrNotDefaultAccount &&
				err != repository.ErrPasswordsDontMatch {
				tools.Logger.Error.Println(err)
				err = errors.New("internal server error")
				code = http.StatusInternalServerError
			}
			return c.Status(code).JSON(fiber.Map{"error": err.Error()})
		}

		return c.SendStatus(http.StatusOK)
	}
}
