package items

import (
	"context"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber"
	"github.com/verryp/gue-eco-test/internal/product/common"
	"github.com/verryp/gue-eco-test/internal/product/consts"
	"github.com/verryp/gue-eco-test/internal/product/handler"
	"github.com/verryp/gue-eco-test/internal/product/presentation"
)

type (
	updateItem struct {
		*handler.Option
	}

	fn func(ctx2 context.Context, id string, r presentation.UpdateItemRequest) (int64, error)
)

func NewUpdateItemHandler(opt *handler.Option) handler.Handler {
	return &updateItem{
		opt,
	}
}

func (h *updateItem) Execute(c *fiber.Ctx) {
	var (
		req      presentation.UpdateItemRequest
		id       = c.Params("id")
		ctx      = c.Context()
		response = common.Response{}
	)

	if req.GrantType == "" {
		req.GrantType = consts.GrantTypeUpdateItemNormal
	}

	err := c.BodyParser(&req)
	if err != nil {
		h.Log.Warn().Msg(err.Error())
		res := response.
			SetStatus(consts.APIStatusError).
			SetMessage(consts.ResponseMessageInvalidRequest)

		c.
			Status(http.StatusBadRequest).
			JSON(res)
		return
	}

	_, err = govalidator.ValidateStruct(req)
	if err != nil {
		h.Log.Warn().Msg(err.Error())

		res := response.
			SetStatus(consts.APIStatusError).
			SetMessage(consts.ResponseMessageInvalidRequest).
			SetErrors(err.Error())

		c.
			Status(http.StatusBadRequest).
			JSON(res)
		return
	}

	updateStatusMap := map[string]fn{
		consts.GrantTypeUpdateItemNormal:   h.Service.Item.UpdateByID,
		consts.GrantTypeUpdateItemDecrease: h.Service.Item.DecreaseItemQuantity,
		consts.GrantTypeUpdateItemIncrease: h.Service.Item.IncreaseItemQuantity,
	}

	rows, err := updateStatusMap[req.GrantType](ctx, id, req)
	if err != nil {
		h.Log.Error().Msg(err.Error())

		res := response.
			SetStatus(consts.APIStatusError).
			SetMessage(consts.ResponseMessageFailedProcessData)

		c.
			Status(http.StatusInternalServerError).
			JSON(res)
		return
	}

	if rows <= 0 {
		res := response.
			SetStatus(consts.APIStatusError).
			SetMessage(consts.ResponseMessageNoDataUpdated)

		c.
			Status(http.StatusUnprocessableEntity).
			JSON(res)
		return
	}

	res := response.
		SetStatus(consts.APIStatusSuccess).
		SetMessage(consts.ResponseMessageSuccessUpdateData)

	c.JSON(res)

	return
}
