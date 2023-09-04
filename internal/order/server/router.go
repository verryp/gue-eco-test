package server

import (
	"github.com/gofiber/fiber"
	"github.com/verryp/gue-eco-test/internal/order/common"
	"github.com/verryp/gue-eco-test/internal/order/consts"
	"github.com/verryp/gue-eco-test/internal/order/handler"
	"github.com/verryp/gue-eco-test/internal/order/handler/aggregator"
	"github.com/verryp/gue-eco-test/internal/order/handler/v1/cart"
	"github.com/verryp/gue-eco-test/internal/order/handler/v1/order"
	updateOrder "github.com/verryp/gue-eco-test/internal/order/handler/v1/order/update"
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

	// carts
	addCart := cart.NewAddCartHandler(rtr.opt)

	// orders
	checkout := order.NewCheckoutHandler(rtr.opt)
	updateOrderAggregator := aggregator.NewUpdateOrderAggregator(rtr.opt, map[string]handler.Handler{
		consts.OrderStatusCanceled: updateOrder.NewCancelHandler(rtr.opt),
	})

	app := fiber.New()

	health := app.Group("health")
	health.Get("/readiness", hcHandler.Readiness)

	v1 := app.Group("v1")

	orders := v1.Group("/orders")
	orders.Post("/checkout", checkout.Execute)
	orders.Put("/:order_id", updateOrderAggregator.Execute)

	carts := v1.Group("/carts")
	carts.Post("/", addCart.Execute)

	return app
}
