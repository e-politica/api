package user

import (
	"github.com/e-politica/api/routes"
	"github.com/e-politica/api/routes/v1/user/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetRoutes(r fiber.Router, tools routes.Tools) {
	registerG := r.Group("/register")

	registerG.Post("/", handlers.PostRegisterDefault(tools))

	loginG := r.Group("/login")
	loginG.Post("/", handlers.PostLoginDefault(tools))
	loginG.Post("/google", handlers.PostLoginGoogle(tools))
}
