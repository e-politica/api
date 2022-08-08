package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/e-politica/api/config"
	v1 "github.com/e-politica/api/routes/v1"
)

func main() {
	app := fiber.New(fiber.Config{})

	// ------------------* Temporary *------------------ //
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("../layout_rascunho/index.html")
	})
	// ------------------* Temporary *------------------ //

	v1.SetRoutes(app.Group("/v1"))

	err := app.Listen(":" + config.ServerPort)
	if err != nil {
		panic(err)
	}
}
