package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/verryp/gue-eco-test/internal/auth/common"
	"github.com/verryp/gue-eco-test/internal/auth/consts"
)

func ValidateClient(c *fiber.Ctx) error {
	response := common.Response{}
	id := c.Get(consts.HeaderClientId)

	if id == "" {
		return c.
			Status(fiber.StatusUnauthorized).
			JSON(response.SetStatus(consts.APIStatusError).SetMessage(consts.ResponseMessageUnauthorized))
	}

	return c.Next()
}
