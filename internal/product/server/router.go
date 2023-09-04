package server

import (
	"github.com/gofiber/fiber"
	"github.com/verryp/gue-eco-test/internal/product/common"
	"github.com/verryp/gue-eco-test/internal/product/handler"
	"github.com/verryp/gue-eco-test/internal/product/handler/v1/items"
)

type (
	router struct {
		config *common.Config
		opt    *handler.Option
		router *fiber.App
	}

	Router interface {
		Route() *fiber.App
	}
)

func NewRouter(cfg *common.Config, opt *handler.Option) Router {
	return &router{
		config: cfg,
		opt:    opt,
		router: fiber.New(&fiber.Settings{
			ReadTimeout:  10,
			WriteTimeout: 10,
		}),
	}
}

func (rtr *router) Route() *fiber.App {
	// health check
	hcHandler := handler.NewHealthCheckHandler(rtr.opt)

	// items
	itemList := items.NewItemListHandler(rtr.opt)
	itemDetail := items.NewItemDetail(rtr.opt)
	itemCreate := items.NewCreateItem(rtr.opt)
	itemUpdate := items.NewUpdateItemHandler(rtr.opt)

	app := fiber.New()

	health := app.Group("health")
	health.Get("/readiness", hcHandler.Readiness)

	v1 := app.Group("v1")

	item := v1.Group("/items")
	item.Get("/", itemList.Execute)
	item.Get("/:id", itemDetail.Execute)
	item.Post("/", itemCreate.Execute)
	item.Put("/:id", itemUpdate.Execute)

	return app
}