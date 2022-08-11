package v1

import (
	"github.com/gofiber/fiber/v2"

	"github.com/e-politica/api/pkg/database"
	"github.com/e-politica/api/routes/v1/user"
)

func SetRoutes(r fiber.Router, db *database.Db) {
	user.SetRoutes(r.Group("/user"), db)
}
