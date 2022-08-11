package handlers

import (
	"log"
	"net/http"

	"github.com/e-politica/api/pkg/database"
	"github.com/e-politica/api/routes/v1/user/repository"
	"github.com/gofiber/fiber/v2"
)

func PostLoginDefault(db *database.Db) fiber.Handler {
	type params struct {
		Credential *string `form:"credential"`
		CsrfToken  *string `form:"g_csrf_token"`
	}

	return func(c *fiber.Ctx) error {
		csrfCookie := c.Request().Header.Cookie("g_csrf_token")
		if csrfCookie == nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Missing 'g_csrf_token' cookie."})
		}

		body := new(params)
		if err := c.BodyParser(body); err != nil {
			log.Println(err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Could not parse request body."})
		}

		if body.Credential == nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Missing 'credential' field in body."})
		}
		if body.CsrfToken == nil {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Missing 'g_csrf_token' field in body."})
		}

		if *body.CsrfToken != string(csrfCookie) {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Failed to verify double submit 'g_csrf_token' cookie."})
		}

		err := repository.Login(c.Context(), *body.Credential)
		if err != nil {
			log.Println(err)
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Error on login."})
		}

		return c.SendStatus(http.StatusOK)
	}
}
