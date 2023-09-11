package order

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/verryp/gue-eco-test/internal/order/common"
	"github.com/verryp/gue-eco-test/internal/order/consts"
	"github.com/verryp/gue-eco-test/internal/order/handler"
)

type list struct {
	*handler.Option
}

func NewOrderListHandler(opt *handler.Option) handler.Handler {
	return &list{
		opt,
	}
}

func (h *list) Execute(c *fiber.Ctx) error {
	var (
		ctx      = c.Context()
		response = common.Response{}
	)

	orders, err := h.Service.Order.List(ctx)
	if err != nil {
		res := response.
			SetStatus(consts.APIStatusError).
			SetMessage(consts.ResponseMessageFailedProcessData)

		return c.
			Status(http.StatusInternalServerError).
			JSON(res)
	}

	if orders == nil {
		return c.SendStatus(http.StatusNoContent)
	}

	res := response.
		SetStatus(consts.APIStatusSuccess).
		SetMessage(consts.ResponseMessageSuccessFetchData).
		SetData(orders)

	return c.JSON(res)
}
