package politician

import (
	"github.com/e-politica/api/routes"
	"github.com/e-politica/api/routes/v1/politician/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetRoutes(r fiber.Router, tools routes.Tools) {
	r.Post("/:id/follow", handlers.PostFollow(tools))
}
