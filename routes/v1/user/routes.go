package user

import (
	"github.com/e-politica/api/routes"
	"github.com/e-politica/api/routes/v1/user/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetRoutes(r fiber.Router, tools routes.Tools) {
	r.Put("/", handlers.PutChangeInfo(tools))
	r.Post("/register", handlers.PostRegisterDefault(tools))
	r.Get("/follows", handlers.GetFollows(tools))

	loginG := r.Group("/login")
	loginG.Post("/", handlers.PostLoginDefault(tools))
	loginG.Post("/google", handlers.PostLoginGoogle(tools))
}
