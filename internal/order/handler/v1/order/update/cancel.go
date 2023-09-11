package update

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/verryp/gue-eco-test/internal/order/common"
	"github.com/verryp/gue-eco-test/internal/order/consts"
	"github.com/verryp/gue-eco-test/internal/order/handler"
	"github.com/verryp/gue-eco-test/internal/order/presentation"
)

type cancel struct {
	*handler.Option
}

func NewCancelHandler(opt *handler.Option) handler.Handler {
	return &cancel{
		opt,
	}
}

func (h *cancel) Execute(c *fiber.Ctx) error {
	var (
		orderID  = c.Params("order_id")
		req      presentation.CancelOrderRequest
		ctx      = c.Context()
		response = common.Response{}
	)

	req.OrderID = orderID

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

	resp, err := h.Service.Order.Cancel(ctx, req)
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
		h.Log.Warn().Msg("the order is not found")

		res := response.
			SetStatus(consts.APIStatusError).
			SetMessage(consts.ResponseMessageDataNotFound)

		return c.
			Status(http.StatusUnprocessableEntity).
			JSON(res)
	}

	res := response.
		SetStatus(consts.APIStatusSuccess).
		SetMessage(consts.ResponseMessageSuccessUpdateData)

	return c.JSON(res)
}
