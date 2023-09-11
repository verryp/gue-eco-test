package service

import (
	"context"

	"github.com/spf13/cast"
	"github.com/verryp/gue-eco-test/internal/order/connector/product"
	"github.com/verryp/gue-eco-test/internal/order/consts"
	"github.com/verryp/gue-eco-test/internal/order/presentation"
	"github.com/verryp/gue-eco-test/internal/order/transformer"
)

type orderSvc struct {
	*Option
}

func NewOrderService(opt *Option) IOrderService {
	return &orderSvc{
		opt,
	}
}

func (svc *orderSvc) Create(ctx context.Context, req presentation.CheckoutRequest) (resp *presentation.CheckoutResponse, err error) {
	ord, err := svc.Repository.Order.FindBy(ctx, req.CartID, consts.OrderStatusOnCart)
	if err != nil {
		svc.Log.Err(err).Msg("error find order")
		return nil, err
	}

	if ord == nil {
		svc.Log.Warn().Msg("order is not found")
		return nil, nil
	}

	detail, err := svc.Repository.OrderDetail.FindByOrderID(ctx, cast.ToString(ord.ID))
	if err != nil {
		svc.Log.Err(err).Msg("error when get order detail")
		return nil, err
	}

	if detail == nil {
		svc.Log.Warn().Msg("order detail is empty")
		return nil, nil
	}

	_, err = svc.ProductAPI.UpdateByID(ctx, cast.ToString(detail.ItemID), product.UpdateProductRequest{
		Quantity:  detail.Quantity,
		GrantType: "decrease",
	})

	if err != nil {
		svc.Log.Err(err).Msg("error update item from product api")
		return nil, err
	}

	p := transformer.ConstructUpdateOrderForCheckout(*ord)
	err = svc.Repository.Order.Update(ctx, p)
	if err != nil {
		svc.Log.Err(err).Msg("err update order")
		return nil, err
	}

	resp = transformer.ToCheckoutResponse(*p, *detail)
	return
}

func (svc *orderSvc) Cancel(ctx context.Context, req presentation.CancelOrderRequest) (*presentation.CheckoutResponse, error) {
	ord, err := svc.Repository.Order.FindBy(ctx, req.OrderID, consts.OrderStatusUnpaid)
	if err != nil {
		return nil, err
	}

	if ord == nil {
		return nil, nil
	}

	detail, err := svc.Repository.OrderDetail.FindByOrderID(ctx, cast.ToString(ord.ID))
	if err != nil {
		svc.Log.Err(err).Msg("error when get order detail")
		return nil, err
	}

	if detail == nil {
		svc.Log.Warn().Msg("order detail is empty")
		return nil, nil
	}

	err = svc.Repository.Order.Update(ctx, transformer.ConstructUpdateOrderForCanceled(req, *ord))
	if err != nil {
		return nil, err
	}

	// re-increase item
	_, err = svc.ProductAPI.UpdateByID(ctx, cast.ToString(detail.ItemID), product.UpdateProductRequest{
		Quantity:  detail.Quantity,
		GrantType: "increase",
	})

	if err != nil {
		svc.Log.Err(err).Msg("error update item from product api")
		return nil, err
	}

	return &presentation.CheckoutResponse{
		ID: cast.ToString(ord.ID),
	}, nil
}

func (svc *orderSvc) List(ctx context.Context) (*presentation.OrderListResponses, error) {
	orders, _, err := svc.Repository.Order.Fetch(ctx)
	if err != nil {
		return nil, err
	}

	if len(orders) == 0 {
		return nil, nil
	}

	return transformer.ToOrderListResponses(orders), nil
}
