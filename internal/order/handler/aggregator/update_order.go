package aggregator

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/verryp/gue-eco-test/internal/order/common"
	"github.com/verryp/gue-eco-test/internal/order/consts"
	"github.com/verryp/gue-eco-test/internal/order/handler"
	"github.com/verryp/gue-eco-test/internal/order/presentation"
)

type updateOrder struct {
	*handler.Option
	updateOrderMap map[string]handler.Handler
}

func NewUpdateOrderAggregator(opt *handler.Option, updateOrderMap map[string]handler.Handler) handler.Handler {
	return &updateOrder{
		Option:         opt,
		updateOrderMap: updateOrderMap,
	}
}

func (h *updateOrder) Execute(c *fiber.Ctx) error {
	var (
		req      presentation.CancelOrderRequest
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

	_, ok := h.updateOrderMap[req.Status]
	if !ok {
		h.Log.Warn().Msg("order status is not supported")

		res := response.
			SetStatus(consts.APIStatusError).
			SetMessage(consts.ResponseMessageInvalidRequest)

		return c.
			Status(http.StatusBadRequest).
			JSON(res)
	}

	return h.updateOrderMap[req.Status].Execute(c)
}
