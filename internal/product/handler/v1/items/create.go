package items

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber"
	"github.com/verryp/gue-eco-test/internal/product/common"
	"github.com/verryp/gue-eco-test/internal/product/consts"
	"github.com/verryp/gue-eco-test/internal/product/handler"
	"github.com/verryp/gue-eco-test/internal/product/presentation"
)

type createItem struct {
	*handler.Option
}

func NewCreateItem(opt *handler.Option) handler.Handler {
	return &createItem{
		Option: opt,
	}
}

func (h *createItem) Execute(c *fiber.Ctx) {
	var (
		req      presentation.CreateItemRequest
		ctx      = c.Context()
		response = common.Response{}
	)

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

	err = h.Service.Item.Create(ctx, req)
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

	res := response.
		SetStatus(consts.APIStatusSuccess).
		SetMessage(consts.ResponseMessageSuccessCreateData)

	c.JSON(res)

	return
}

// func (h *createItem) validate(req presentation.CreateItemRequest) url.Values {
// 	v := validator.New()
// 	errs := v.Struct(&req)
// 	var ev url.Values
//
// 	if errs != nil {
// 		for _, err := range errs.(validator.ValidationErrors) {
// 			ev.Add(err.Field(), fmt.Sprintf("is %s", err.Value()))
// 		}
// 	}
//
// 	return ev
// }
