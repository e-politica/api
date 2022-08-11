package user

import (
	"github.com/e-politica/api/pkg/database"
	"github.com/e-politica/api/routes/v1/user/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetRoutes(r fiber.Router, db *database.Db) {
	loginG := r.Group("/login")
	loginG.Post("/", handlers.PostLoginDefault(db))
	loginG.Post("/google", handlers.PostLoginGoogle(db))

	registerG := r.Group("/register")
	registerG.Post("/", handlers.PostRegisterDefault(db))
}
