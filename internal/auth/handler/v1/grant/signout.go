package grant

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/verryp/gue-eco-test/internal/auth/common"
	"github.com/verryp/gue-eco-test/internal/auth/consts"
	"github.com/verryp/gue-eco-test/internal/auth/handler"
	"github.com/verryp/gue-eco-test/internal/auth/helper"
)

type signOut struct {
	*handler.Option
}

func NewSignOutHandler(opt *handler.Option) handler.Handler {
	return &signOut{
		opt,
	}
}

func (h *signOut) Execute(c *fiber.Ctx) error {
	var (
		token    = c.Get("authorization")
		ctx      = c.Context()
		response = common.Response{}
	)

	err := h.Service.Auth.BlackListToken(ctx, helper.SplitHeaderBearerToken(token))
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
		SetMessage(consts.ResponseMessageSuccessSignOut)

	return c.JSON(res)
}
