package grant

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/verryp/gue-eco-test/internal/auth/common"
	"github.com/verryp/gue-eco-test/internal/auth/consts"
	"github.com/verryp/gue-eco-test/internal/auth/handler"
	"github.com/verryp/gue-eco-test/internal/auth/presentation"
)

type reToken struct {
	*handler.Option
}

func NewReTokenHandler(opt *handler.Option) handler.Handler {
	return &reToken{
		opt,
	}
}

func (h *reToken) Execute(c *fiber.Ctx) error {
	var (
		token    = c.Get("authorization")
		clientID = c.Get("X-Client-Id")
		path     = string(c.Request().URI().Path())
		ctx      = c.Context()
		response = common.Response{}
	)

	resp, err := h.Service.Auth.RefreshToken(ctx, presentation.ReTokenRequest{
		Token:    token,
		ClientID: clientID,
		PathURL:  path,
	})
	if err != nil {
		h.Log.Error().Msg(err.Error())

		res := response.
			SetStatus(consts.APIStatusError).
			SetMessage(consts.ResponseMessageFailedProcessData)

		return c.
			Status(http.StatusInternalServerError).
			JSON(res)
	}

	res := response.
		SetStatus(consts.APIStatusSuccess).
		SetMessage(consts.ResponseMessageSuccessRefreshToken).
		SetData(resp)

	return c.JSON(res)
}
