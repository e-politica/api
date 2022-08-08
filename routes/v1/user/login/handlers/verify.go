package handlers

import (
	"log"
	"net/http"

	"github.com/e-politica/api/routes/v1/user/login/repository"
	"github.com/gofiber/fiber/v2"
)

type verifyParams struct {
	State *string `query:"state"`
	Code  *string `query:"code"`
}

func GetVerify(c *fiber.Ctx) error {
	query := new(verifyParams)
	if err := c.QueryParser(query); err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Could not parse request query."})
	}

	if query.State == nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Missing 'state' field in query."})
	}
	if query.Code == nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Missing 'code' field in query."})
	}

	if *query.State != c.Cookies("STATE") {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Query and cookie 'state' doesn't match."})
	}

	if err := repository.Verify(c.Context(), *query.Code); err != nil {
		log.Println(err)
		return c.Status(http.StatusInternalServerError).JSON(fiber.Map{"error": "Could not verify token."})
	}

	return c.SendStatus(http.StatusOK)
}
