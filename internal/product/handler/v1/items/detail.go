package items

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
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

func (h *itemDetail) Execute(c *fiber.Ctx) error {
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

		return c.
			Status(http.StatusInternalServerError).
			JSON(res)
	}

	if item == nil {
		res := response.
			SetStatus(consts.APIStatusError).
			SetMessage(consts.ResponseMessageDataNotFound)

		return c.
			Status(http.StatusUnprocessableEntity).
			JSON(res)
	}

	res := response.
		SetStatus(consts.APIStatusSuccess).
		SetMessage(consts.ResponseMessageSuccessFoundData).
		SetData(item)

	return c.JSON(res)
}
