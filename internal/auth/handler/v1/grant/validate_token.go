package grant

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/verryp/gue-eco-test/internal/auth/common"
	"github.com/verryp/gue-eco-test/internal/auth/consts"
	"github.com/verryp/gue-eco-test/internal/auth/handler"
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

	token := c.Get("authorization")

	resp, err := h.Service.Auth.ValidateToken(ctx, token)
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
