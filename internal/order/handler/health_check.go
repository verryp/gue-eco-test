package handler

import (
	"net/http"

	"github.com/gofiber/fiber"
)

type healthCheckHandler struct {
	*Option
}

func NewHealthCheckHandler(opt *Option) *healthCheckHandler {
	return &healthCheckHandler{opt}
}

func (h *healthCheckHandler) Readiness(c *fiber.Ctx) {
	ctx := c.Context()

	err := h.Service.HealthCheck.HealthCheck(ctx)
	if err != nil {
		c.
			Status(http.StatusInternalServerError).
			Send(err.Error())
		return
	}

	res := struct {
		HealthCheck string `json:"status"`
	}{HealthCheck: "ok"}

	c.JSON(res)
}
