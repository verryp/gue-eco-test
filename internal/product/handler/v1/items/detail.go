package items

import (
	"net/http"

	"github.com/gofiber/fiber"
	"github.com/verryp/gue-eco-test/internal/product/common"
	"github.com/verryp/gue-eco-test/internal/product/consts"
	"github.com/verryp/gue-eco-test/internal/product/handler"
)

type itemDetail struct {
	*handler.Option
}

func NewItemDetail(opt *handler.Option) handler.Handler {
	return &itemDetail{
		opt,
	}
}

func (h *itemDetail) Execute(c *fiber.Ctx) {
	var (
		id       = c.Params("id")
		ctx      = c.Context()
		response = common.Response{}
	)

	item, err := h.Service.Item.FindByID(ctx, id)
	if err != nil {
		res := response.
			SetStatus(consts.APIStatusError).
			SetMessage(consts.ResponseMessageFailedProcessData)

		c.
			Status(http.StatusInternalServerError).
			JSON(res)
		return
	}

	if item == nil {
		res := response.
			SetStatus(consts.APIStatusError).
			SetMessage(consts.ResponseMessageDataNotFound)

		c.
			Status(http.StatusUnprocessableEntity).
			JSON(res)

		return
	}

	res := response.
		SetStatus(consts.APIStatusSuccess).
		SetMessage(consts.ResponseMessageSuccessFoundData).
		SetData(item)

	c.JSON(res)

	return
}
