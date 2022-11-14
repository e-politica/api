package handlers

import (
	"errors"
	"net/http"

	"github.com/e-politica/api/routes"
	"github.com/e-politica/api/routes/v1/user/repository"
	"github.com/gofiber/fiber/v2"
)

func GetPublicInfo(tools routes.Tools) fiber.Handler {
	return func(c *fiber.Ctx) error {
		id := c.Params("id")
		if id == "" {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "must provide url param 'id'"})
		}

		info, err := repository.GetPublicInfo(c.Context(), tools.Db, id)
		if err != nil {
			tools.Logger.Error.Println(err)
			err = errors.New("internal server error")
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
		}

		return c.Status(http.StatusOK).JSON(info)
	}
}
