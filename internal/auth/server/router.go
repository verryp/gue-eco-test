package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/verryp/gue-eco-test/internal/auth/common"
	"github.com/verryp/gue-eco-test/internal/auth/handler"
	"github.com/verryp/gue-eco-test/internal/auth/handler/v1/grant"
	"github.com/verryp/gue-eco-test/internal/auth/handler/v1/register"
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
		router: fiber.New(fiber.Config{
			ReadTimeout:  10,
			WriteTimeout: 10,
		}),
	}
}

func (rtr *router) Route() *fiber.App {
	// health check
	hcHandler := handler.NewHealthCheckHandler(rtr.opt)

	// auth
	signup := register.NewSignupHandler(rtr.opt)
	authorizeClient := grant.NewClientAuthorization(rtr.opt)
	validateToken := grant.NewValidateTokenHandler(rtr.opt)
	signin := grant.NewSignInHandler(rtr.opt)

	app := fiber.New()

	health := app.Group("health")
	health.Get("/readiness", hcHandler.Readiness)

	v1 := app.Group("v1")

	auth := v1.Group("/auth")
	auth.Post("/signup", signup.Execute)

	auth.Post("/authorize", authorizeClient.Execute)
	auth.Post("/token", validateToken.Execute)
	auth.Post("/signin", func(c *fiber.Ctx) error {
		clientId := c.Get("X-Client-Id")

		if clientId == "" {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		return c.Next()
	}, signin.Execute)

	return app
}
