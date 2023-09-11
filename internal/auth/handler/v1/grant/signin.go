package grant

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/verryp/gue-eco-test/internal/auth/common"
	"github.com/verryp/gue-eco-test/internal/auth/consts"
	"github.com/verryp/gue-eco-test/internal/auth/handler"
	"github.com/verryp/gue-eco-test/internal/auth/presentation"
)

type signIn struct {
	*handler.Option
}

func NewSignInHandler(opt *handler.Option) handler.Handler {
	return &signIn{
		opt,
	}
}

func (h *signIn) Execute(c *fiber.Ctx) error {
	var (
		req      presentation.SignInRequest
		path     = string(c.Request().URI().Path())
		ctx      = c.Context()
		response = common.Response{}
	)

	req.ClientID = c.Get("X-Client-Id")

	err := c.BodyParser(&req)
	if err != nil {
		h.Log.Warn().Msg(err.Error())
		res := response.
			SetStatus(consts.APIStatusError).
			SetMessage(consts.ResponseMessageInvalidRequest)

		return c.
			Status(http.StatusBadRequest).
			JSON(res)
	}

	_, err = govalidator.ValidateStruct(req)
	if err != nil {
		h.Log.Warn().Msg(err.Error())

		res := response.
			SetStatus(consts.APIStatusError).
			SetMessage(consts.ResponseMessageInvalidRequest).
			SetErrors(err.Error())

		return c.
			Status(http.StatusBadRequest).
			JSON(res)
	}

	resp, err := h.Service.Auth.Login(ctx, path, req)
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
		SetMessage(consts.ResponseMessageSuccessSignIn).
		SetData(resp)

	return c.JSON(res)
}
