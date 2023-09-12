package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/verryp/gue-eco-test/internal/auth/common"
	"github.com/verryp/gue-eco-test/internal/auth/handler"
	"github.com/verryp/gue-eco-test/internal/auth/handler/v1/grant"
	"github.com/verryp/gue-eco-test/internal/auth/handler/v1/register"
	"github.com/verryp/gue-eco-test/internal/auth/middleware"
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
	signIn := grant.NewSignInHandler(rtr.opt)
	reToken := grant.NewReTokenHandler(rtr.opt)
	signOut := grant.NewSignOutHandler(rtr.opt)

	app := fiber.New()

	health := app.Group("health")
	health.Get("/readiness", hcHandler.Readiness)

	v1 := app.Group("v1")

	auth := v1.Group("/auth")

	auth.Post("/signup", middleware.ValidateClient, signup.Execute)
	auth.Post("/authorize", authorizeClient.Execute)
	auth.Post("/token", validateToken.Execute)
	auth.Post("/signin", middleware.ValidateClient, signIn.Execute)
	auth.Post("/retoken", middleware.ValidateUser, reToken.Execute)
	auth.Post("/signout", middleware.ValidateUser, signOut.Execute)

	return app
}
