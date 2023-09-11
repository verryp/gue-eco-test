package product

import (
	"context"
)

type (
	API interface {
		GetByID(ctx context.Context, id string) (*GetProductResponse, error)
		UpdateByID(ctx context.Context, id string, req UpdateProductRequest) (*BaseResponse, error)
	}
)
