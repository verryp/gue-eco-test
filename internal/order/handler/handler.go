package handler

import (
	"github.com/gofiber/fiber"
	"github.com/verryp/gue-eco-test/internal/order/common"
	"github.com/verryp/gue-eco-test/internal/order/service"
)

type (
	Option struct {
		*common.Option
		Service *service.Service
	}

	Handler interface {
		Execute(c *fiber.Ctx)
	}
)
