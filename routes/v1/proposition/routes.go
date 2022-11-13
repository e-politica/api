package proposition

import (
	"github.com/e-politica/api/routes"
	"github.com/e-politica/api/routes/v1/proposition/handlers"
	"github.com/gofiber/fiber/v2"
)

func SetRoutes(r fiber.Router, tools routes.Tools) {
	r.Post("/:id/like", handlers.PostLike(tools))
	r.Post("/:id/comment", handlers.PostComment(tools))
	r.Get("/:id/comments", handlers.GetComments(tools))
}
