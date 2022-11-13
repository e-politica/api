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

func PostLoginGoogle(tools routes.Tools) fiber.Handler {
	return func(c *fiber.Ctx) error {
		csrfCookie := c.Request().Header.Cookie("g_csrf_token")
		if csrfCookie == nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "missing 'g_csrf_token' cookie"})
		}

		var params user.LoginGoogleParams
		if err := c.BodyParser(&params); err != nil {
			tools.Logger.Error.Println(err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "could not parse request body"})
		}

		if err := params.Validate(csrfCookie); err != nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
		}

		userSession, err := repository.LoginGoogle(c.Context(), tools.Db, params.Credential)
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
