package handler

import (
	"github.com/gofiber/fiber/v2"
	"github.com/verryp/gue-eco-test/internal/auth/common"
	"github.com/verryp/gue-eco-test/internal/auth/service"
)

type (
	Option struct {
		*common.Option
		Service *service.Service
	}

	Handler interface {
		Execute(c *fiber.Ctx) error
	}
)
