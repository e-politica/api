package login

import (
	"github.com/gofiber/fiber/v2"

	"github.com/e-politica/api/routes/v1/user/login/handlers"
)

func SetRoutes(r fiber.Router) {
	r.Post("/", handlers.PostLogin)
	r.Get("/verify", handlers.GetVerify)
}
