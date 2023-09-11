package product

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/verryp/gue-eco-test/internal/order/common"
	"github.com/verryp/gue-eco-test/pkg/httpclient"
)

type productAPI struct {
	*common.Config
	*httpclient.RestClient
}

func NewProductAPI(opt *common.Config, restClient *httpclient.RestClient) API {
	return &productAPI{
		Config:     opt,
		RestClient: restClient,
	}
}

func (api *productAPI) GetByID(ctx context.Context, id string) (*GetProductResponse, error) {
	res, err := api.RestClient.R().
		SetHeaders(map[string]string{
			"Content-Type": "application/json",
		}).
		Get(fmt.Sprintf("%s/%s", api.Config.Dependency.Product.GetDetailItemPath, id))

	if err != nil {
		return nil, err
	}

	if !res.IsSuccess() {
		return nil, fmt.Errorf(res.String())
	}

	var result GetProductResponse
	err = json.Unmarshal(res.Body(), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func (api *productAPI) UpdateByID(ctx context.Context, id string, req UpdateProductRequest) (*BaseResponse, error) {
	res, err := api.RestClient.R().
		EnableTrace().
		SetBody(req).
		SetHeader("Content-Type", "application/json").
		Put(fmt.Sprintf("%s/%s", api.Config.Dependency.Product.UpdateItemPath, id))

	if err != nil {
		return nil, err
	}

	if !res.IsSuccess() {
		return nil, fmt.Errorf(res.String())
	}

	var result BaseResponse
	err = json.Unmarshal(res.Body(), &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}
