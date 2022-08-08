package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/e-politica/api/routes/v1/user/login/repository"
	"github.com/gofiber/fiber/v2"
)

type loginParams struct {
	Credential *string `form:"credential"`
	CsrfToken  *string `form:"g_csrf_token"`
}

func PostLogin(c *fiber.Ctx) error {

	csrfCookie := c.Request().Header.Cookie("g_csrf_token")
	if csrfCookie == nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Missing 'g_csrf_token' cookie."})
	}

	body := new(loginParams)
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

	state, url := repository.Login()
	c.Cookie(&fiber.Cookie{
		Name:    "STATE",
		Value:   state,
		Expires: time.Now().Add(time.Minute),
	})
	return c.Redirect(url)
}
