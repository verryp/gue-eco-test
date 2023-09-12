package grant

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/verryp/gue-eco-test/internal/auth/common"
	"github.com/verryp/gue-eco-test/internal/auth/consts"
	"github.com/verryp/gue-eco-test/internal/auth/handler"
	"github.com/verryp/gue-eco-test/internal/auth/helper"
	"github.com/verryp/gue-eco-test/internal/auth/presentation"
)

type validateToken struct {
	*handler.Option
}

func NewValidateTokenHandler(opt *handler.Option) handler.Handler {
	return &validateToken{
		Option: opt,
	}
}

func (h *validateToken) Execute(c *fiber.Ctx) error {
	var (
		ctx      = c.Context()
		response = common.Response{}
	)

	token := c.Get(consts.HeaderAuthorization)
	forwardedFor := c.Get(consts.HeaderForwardedFor)
	userAgent := c.Get(consts.HeaderUserAgent)
	pathActivity := c.Get(consts.HeaderPathSource)

	resp, err := h.Service.Auth.ValidateToken(ctx, presentation.ValidateTokenRequest{
		IPAddress: helper.TrimHeaderIPAddress(forwardedFor),
		UserAgent: userAgent,
		Path:      pathActivity,
		Token:     token,
	})
	if err != nil {
		h.Log.Error().Msg(err.Error())

		res := response.
			SetStatus(consts.APIStatusError).
			SetMessage(consts.ResponseMessageUnauthorized)

		return c.
			Status(http.StatusUnauthorized).
			JSON(res)
	}

	res := response.
		SetStatus(consts.APIStatusSuccess).
		SetMessage(consts.ResponseMessageTokenValidated).
		SetData(resp)

	return c.JSON(res)
}
