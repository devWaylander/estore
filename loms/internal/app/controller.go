package orders

import (
	"context"
	"log"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	servicepb "route256.ozon.ru/project/loms/gen/api/orders/v1"
	internalErrors "route256.ozon.ru/project/loms/internal/errors"
	"route256.ozon.ru/project/loms/internal/model"
)

var _ servicepb.LOMSServer = (*controller)(nil)

type LOMSService interface {
	OrderCreate(ctx context.Context, data model.Order) (uint64, error)
	OrderGetInfo(ctx context.Context, orderID uint64) (model.Info, error)
	OrderPay(ctx context.Context, orderID uint64) error
	OrderCancel(ctx context.Context, orderID uint64) error
	OrderGetStockInfo(ctx context.Context, SKU uint32) (uint64, error)
}

type controller struct {
	servicepb.UnimplementedLOMSServer
	service LOMSService
}

func New(service LOMSService) *controller {
	return &controller{service: service}
}

func (c *controller) OrderCreate(ctx context.Context, data *servicepb.OrderCreateReq) (*servicepb.OrderCreateResp, error) {
	id, err := c.service.OrderCreate(ctx, model.ToOrder(data.Order))
	if err != nil {
		if err.Error() == internalErrors.ErrStocksAreEmpty {
			return &servicepb.OrderCreateResp{OrderId: 0}, status.Error(codes.FailedPrecondition, err.Error())
		}
		log.Println(err)
	}

	return &servicepb.OrderCreateResp{OrderId: id}, nil
}

func (c *controller) OrderGetInfo(ctx context.Context, data *servicepb.OrderGetInfoReq) (*servicepb.OrderGetInfoResp, error) {
	info, err := c.service.OrderGetInfo(ctx, data.OrderId)
	if err != nil {
		if err.Error() == internalErrors.ErrOrderNotFound {
			return &servicepb.OrderGetInfoResp{}, status.Error(codes.NotFound, err.Error())
		}
		log.Println(err)
	}

	return model.ToOrderGetInfoRespServicepb(info), nil
}

func (c *controller) OrderPay(ctx context.Context, data *servicepb.OrderPayReq) (*servicepb.OrderPayResp, error) {
	err := c.service.OrderPay(ctx, data.OrderId)
	if err != nil {
		log.Println(err)
	}

	return &servicepb.OrderPayResp{}, nil
}

func (c *controller) OrderCancel(ctx context.Context, data *servicepb.OrderCancelReq) (*servicepb.OrderCancelResp, error) {
	err := c.service.OrderCancel(ctx, data.OrderId)
	if err != nil {
		log.Println()
	}

	return &servicepb.OrderCancelResp{}, nil
}

func (c *controller) OrderGetStockInfo(ctx context.Context, data *servicepb.OrderGetStockInfoReq) (*servicepb.OrderGetStockInfoResp, error) {
	count, err := c.service.OrderGetStockInfo(ctx, data.Sku)
	if err != nil {
		log.Println()
	}

	return &servicepb.OrderGetStockInfoResp{Count: count}, nil
}
