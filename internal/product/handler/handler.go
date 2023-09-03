package handler

import (
	"github.com/gofiber/fiber"
	"github.com/verryp/gue-eco-test/internal/product/common"
	"github.com/verryp/gue-eco-test/internal/product/service"
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
