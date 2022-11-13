package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/e-politica/api/pkg/session"
	"github.com/e-politica/api/routes"
	"github.com/e-politica/api/routes/v1/user/repository"
	"github.com/gofiber/fiber/v2"
)

func GetFollows(tools routes.Tools) fiber.Handler {
	return func(c *fiber.Ctx) error {
		access := c.Get("Authorization")
		if access == "" {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "must provide 'Authorization' header"})
		}

		rawPage := c.Query("page", "1")
		page, err := strconv.Atoi(rawPage)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "must provide a valid url query 'page'"})
		}

		page--
		if page < 0 {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "'page' url query must be >= 1"})
		}

		rawLimit := c.Query("limit", "15")
		limit, err := strconv.Atoi(rawLimit)
		if err != nil {
			return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "must provide a valid url param 'limit'"})
		}

		follows, err := repository.GetFollows(c.Context(), tools.Db, access, page*limit, limit)
		if err != nil {
			code := http.StatusBadRequest
			if err != session.ErrSessionNotFound {
				tools.Logger.Error.Println(err)
				err = errors.New("internal server error")
				code = http.StatusInternalServerError
			}
			return c.Status(code).JSON(fiber.Map{"error": err.Error()})
		}

		return c.Status(http.StatusOK).JSON(follows)
	}
}
