package v1

import (
	"github.com/e-politica/api/routes"

	"github.com/gofiber/fiber/v2"

	"github.com/e-politica/api/routes/v1/user"
)

func SetRoutes(r fiber.Router, tools routes.Tools) {
	user.SetRoutes(r.Group("/user"), tools)
}
