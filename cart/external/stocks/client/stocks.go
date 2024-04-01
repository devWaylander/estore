package stocks

import (
	"context"
	"errors"

	desc "route256.ozon.ru/project/cart/external/stocks/gen/api/orders/v1"
	internalErrors "route256.ozon.ru/project/cart/internal/errors"
)

type client struct {
	header string
	client desc.LOMSClient
}

func New(header string, grpcClient desc.LOMSClient) *client {
	return &client{
		header: header,
		client: grpcClient,
	}
}

func (c *client) GetStocksInfo(ctx context.Context, SKU uint32) (uint64, error) {
	resp, err := c.client.OrderGetStockInfo(ctx, &desc.OrderGetStockInfoReq{Sku: SKU})
	if err != nil {
		return 0, errors.New(internalErrors.ErrStockCount)
	}

	return resp.Count, nil
}

func (c *client) CreateOrder(ctx context.Context, order *desc.Order) (uint64, error) {
	resp, err := c.client.OrderCreate(ctx, &desc.OrderCreateReq{Order: order})
	if err != nil {
		return 0, errors.New(internalErrors.ErrOrder)
	}

	return resp.OrderId, nil
}
