package grant

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/verryp/gue-eco-test/internal/auth/common"
	"github.com/verryp/gue-eco-test/internal/auth/consts"
	"github.com/verryp/gue-eco-test/internal/auth/handler"
)

type clientAuthorization struct {
	*handler.Option
}

func NewClientAuthorization(opt *handler.Option) handler.Handler {
	return &clientAuthorization{
		Option: opt,
	}
}

func (h *clientAuthorization) Execute(c *fiber.Ctx) error {
	var (
		ctx      = c.Context()
		response = common.Response{}
	)

	// !note: this is must be by db, but simplify with hardcode
	apiKey := c.Get("authorization")

	fmt.Println("apikey", apiKey)

	token, err := h.Service.Auth.ClientAuthorization(ctx, apiKey, string(c.Request().URI().Path()))
	if err != nil {
		h.Log.Error().Msg(err.Error())

		res := response.
			SetStatus(consts.APIStatusError).
			SetMessage(consts.ResponseMessageUnauthorized)

		return c.
			Status(http.StatusUnauthorized).
			JSON(res)
	}

	if token == nil {
		h.Log.Error().Msg("api key is invalid")

		res := response.
			SetStatus(consts.APIStatusError).
			SetMessage(consts.ResponseMessageUnauthorized)

		return c.
			Status(http.StatusUnauthorized).
			JSON(res)
	}

	res := response.
		SetStatus(consts.APIStatusSuccess).
		SetMessage(consts.ResponseMessageSuccessAuthorizedClient).
		SetData(token)

	return c.JSON(res)
}
