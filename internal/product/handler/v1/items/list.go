package items

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/verryp/gue-eco-test/internal/product/common"
	"github.com/verryp/gue-eco-test/internal/product/consts"
	"github.com/verryp/gue-eco-test/internal/product/handler"
)

type itemListHandler struct {
	*handler.Option
}

func NewItemListHandler(opt *handler.Option) handler.Handler {
	return &itemListHandler{
		Option: opt,
	}
}

func (h *itemListHandler) Execute(c *fiber.Ctx) error {
	var (
		ctx      = c.Context()
		response = common.Response{}
	)

	items, err := h.Service.Item.List(ctx)
	if err != nil {
		res := response.
			SetStatus(consts.APIStatusError).
			SetMessage(consts.ResponseMessageFailedProcessData)

		return c.
			Status(http.StatusInternalServerError).
			JSON(res)

	}

	if len(items.Items) == 0 {
		return c.SendStatus(http.StatusNoContent)
	}

	res := response.
		SetStatus(consts.APIStatusSuccess).
		SetMessage(consts.ResponseMessageSuccessFetchData).
		SetData(items)

	return c.JSON(res)
}
