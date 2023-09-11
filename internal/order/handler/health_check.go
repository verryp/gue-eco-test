package handler

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

type healthCheckHandler struct {
	*Option
}

func NewHealthCheckHandler(opt *Option) *healthCheckHandler {
	return &healthCheckHandler{opt}
}

func (h *healthCheckHandler) Readiness(c *fiber.Ctx) error {
	ctx := c.Context()

	err := h.Service.HealthCheck.HealthCheck(ctx)
	if err != nil {
		return c.
			Status(http.StatusInternalServerError).
			SendString(err.Error())
	}

	res := struct {
		HealthCheck string `json:"status"`
	}{HealthCheck: "ok"}

	return c.JSON(res)
}
