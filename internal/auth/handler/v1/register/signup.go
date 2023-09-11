package register

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/verryp/gue-eco-test/internal/auth/common"
	"github.com/verryp/gue-eco-test/internal/auth/consts"
	"github.com/verryp/gue-eco-test/internal/auth/handler"
	"github.com/verryp/gue-eco-test/internal/auth/presentation"
)

type signup struct {
	*handler.Option
}

func NewSignupHandler(opt *handler.Option) handler.Handler {
	return &signup{
		opt,
	}
}

func (h *signup) Execute(c *fiber.Ctx) error {
	var (
		req      presentation.SignUpRequest
		ctx      = c.Context()
		response = common.Response{}
	)

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

	err = h.Service.Auth.Register(ctx, req)
	if err != nil {
		h.Log.Error().Msg(err.Error())

		res := response.
			SetStatus(consts.APIStatusError).
			SetMessage(consts.ResponseMessageFailedProcessData)

		return c.
			Status(http.StatusUnprocessableEntity).
			JSON(res)
	}

	res := response.
		SetStatus(consts.APIStatusSuccess).
		SetMessage(consts.ResponseMessageSuccessSignup)

	return c.JSON(res)
}
