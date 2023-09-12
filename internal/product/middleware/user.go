package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/verryp/gue-eco-test/internal/product/common"
	"github.com/verryp/gue-eco-test/internal/product/consts"
)

func ValidateUser(c *fiber.Ctx) error {
	response := common.Response{}
	email := c.Get(consts.HeaderUserEmail)
	id := c.Get(consts.HeaderUserID)
	name := c.Get(consts.HeaderUserName)

	if id == "" || email == "" || name == "" {
		return c.
			Status(fiber.StatusUnauthorized).
			JSON(response.SetStatus(consts.APIStatusError).SetMessage(consts.ResponseMessageUnauthorized))
	}

	return c.Next()
}
