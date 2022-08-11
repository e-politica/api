package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/e-politica/api/config"
	"github.com/e-politica/api/pkg/database"
	v1 "github.com/e-politica/api/routes/v1"
)

func main() {
	db := database.New()
	defer db.Conn.Close(*db.Ctx)
	go db.LoopCheckConnection()

	app := fiber.New(fiber.Config{})

	// ------------------* Temporary *------------------ //
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("../layout_rascunho/index.html")
	})
	// ------------------* Temporary *------------------ //

	v1.SetRoutes(app.Group("/v1"), db)

	err := app.Listen(":" + config.ServerPort)
	if err != nil {
		panic(err)
	}
}
