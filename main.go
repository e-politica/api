package main

import (
	"context"
	"os"

	"github.com/gofiber/fiber/v2"

	"github.com/e-politica/api/config"
	"github.com/e-politica/api/pkg/database"
	"github.com/e-politica/api/pkg/log"
	"github.com/e-politica/api/routes"
	v1 "github.com/e-politica/api/routes/v1"
)

func main() {
	ctx := context.Background()
	db := database.New(&ctx)
	defer db.Conn.Close(*db.Ctx)
	go db.LoopCheckConnection()

	tools := routes.Tools{
		Db:     db,
		Logger: log.NewLogger(os.Stdout),
	}

	app := fiber.New(fiber.Config{})

	// ------------------* Temporary *------------------ //
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendFile("../layout_rascunho/rascunho2/login.html")
	})
	// ------------------* Temporary *------------------ //

	v1.SetRoutes(app.Group("/v1"), tools)

	err := app.Listen(":" + config.ServerPort)
	if err != nil {
		panic(err)
	}
}
