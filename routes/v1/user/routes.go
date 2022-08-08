package user

import (
	"github.com/gofiber/fiber/v2"

	"github.com/e-politica/api/routes/v1/user/login"
)

func SetRoutes(r fiber.Router) {
	login.SetRoutes(r.Group("/login"))
}
