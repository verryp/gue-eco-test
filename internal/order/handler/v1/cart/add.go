package cart

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/verryp/gue-eco-test/internal/order/common"
	"github.com/verryp/gue-eco-test/internal/order/consts"
	"github.com/verryp/gue-eco-test/internal/order/handler"
	"github.com/verryp/gue-eco-test/internal/order/presentation"
)

type addCart struct {
	*handler.Option
}

func NewAddCartHandler(opt *handler.Option) handler.Handler {
	return &addCart{
		opt,
	}
}

func (h *addCart) Execute(c *fiber.Ctx) error {
	var (
		req      presentation.AddCartRequest
		ctx      = c.Context()
		response = common.Response{}
	)

	req.CustomerName = c.Get(consts.HeaderUserName)
	req.CustomerEmail = c.Get(consts.HeaderUserEmail)

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

	err = h.Service.Cart.Add(ctx, req)
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
		SetMessage(consts.ResponseMessageSuccessCreateData)

	return c.JSON(res)
}
