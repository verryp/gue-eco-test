package order

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/verryp/gue-eco-test/internal/order/common"
	"github.com/verryp/gue-eco-test/internal/order/consts"
	"github.com/verryp/gue-eco-test/internal/order/handler"
	"github.com/verryp/gue-eco-test/internal/order/presentation"
)

type checkout struct {
	*handler.Option
}

func NewCheckoutHandler(opt *handler.Option) handler.Handler {
	return &checkout{
		opt,
	}
}

func (h *checkout) Execute(c *fiber.Ctx) error {
	var (
		req      presentation.CheckoutRequest
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

	resp, err := h.Service.Order.Create(ctx, req)
	if err != nil {
		h.Log.Error().Msg(err.Error())

		res := response.
			SetStatus(consts.APIStatusError).
			SetMessage(consts.ResponseMessageFailedProcessData)

		return c.
			Status(http.StatusInternalServerError).
			JSON(res)
	}

	if resp == nil {
		h.Log.Warn().Msg("order failed created")

		res := response.
			SetStatus(consts.APIStatusError).
			SetMessage(consts.ResponseMessageOrderFailedCreated)

		return c.
			Status(http.StatusUnprocessableEntity).
			JSON(res)
	}

	res := response.
		SetStatus(consts.APIStatusSuccess).
		SetMessage(consts.ResponseMessageSuccessUpdateData).
		SetData(resp)

	return c.JSON(res)
}
