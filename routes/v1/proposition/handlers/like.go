package handlers

import (
	"errors"
	"net/http"

	"github.com/e-politica/api/routes"
	"github.com/e-politica/api/routes/v1/proposition/repository"
	"github.com/gofiber/fiber/v2"
)

func PostLike(tools routes.Tools) fiber.Handler {
	return func(c *fiber.Ctx) error {
		access := c.Get("Authorization")
		if access == "" {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "must provide 'Authorization' header"})
		}

		id := c.Params("id")
		if id == "" {
			return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "must provide url param 'id'"})
		}

		err := repository.Like(c.Context(), tools.Db, access, id)
		if err != nil {
			code := http.StatusBadRequest
			if err != repository.ErrAlreadySigned &&
				err != repository.ErrPropositionNotFound {
				tools.Logger.Error.Println(err)
				err = errors.New("internal server error")
				code = http.StatusInternalServerError
			}
			return c.Status(code).JSON(fiber.Map{"error": err.Error()})
		}

		return c.SendStatus(http.StatusOK)
	}
}
