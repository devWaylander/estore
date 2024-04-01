package service

import (
	"context"
	"errors"
	"log"

	internalErrors "route256.ozon.ru/project/loms/internal/errors"
	"route256.ozon.ru/project/loms/internal/model"
)

type Repository interface {
	CreateOrder(ctx context.Context, data model.Order) (uint64, error)
	GetOrder(ctx context.Context, orderID uint64) (model.Order, error)
	UpdateOrderStatus(ctx context.Context, orderID uint64, status string) error
	GetStockBySKU(ctx context.Context, SKU uint32) model.Stock
	UpdateStock(ctx context.Context, data model.Stock) error
}

type service struct {
	repo Repository
}

func New(repo Repository) *service {
	return &service{
		repo: repo,
	}
}

func (s *service) updateOrderStatus(ctx context.Context, data model.Order) error {
	err := s.repo.UpdateOrderStatus(ctx, data.ID, data.Status)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *service) stocksReserve(ctx context.Context, data []model.Item) error {
	for _, e := range data {
		stock := s.repo.GetStockBySKU(ctx, e.SKU)
		if e.Count > (stock.TotalCount - stock.Reserved) {
			return errors.New(internalErrors.ErrStocksAreEmpty)
		}

		stock.Reserved = stock.Reserved + e.Count
		err := s.repo.UpdateStock(ctx, stock)
		if err != nil {
			log.Println(err)
			return err
		}
	}

	return nil
}

func (s *service) stocksRemove(ctx context.Context, data []model.Item, isCancel bool) error {
	for _, e := range data {
		stock := s.repo.GetStockBySKU(ctx, e.SKU)
		stock.Reserved = stock.Reserved - e.Count

		if !isCancel {
			stock.TotalCount = stock.TotalCount - e.Count
		}

		err := s.repo.UpdateStock(ctx, stock)
		if err != nil {
			log.Println(err)
			return err
		}
	}

	return nil
}

func (s *service) OrderCreate(ctx context.Context, data model.Order) (uint64, error) {
	data.Status = model.StatusNew
	id, err := s.repo.CreateOrder(ctx, data)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	err = s.stocksReserve(ctx, data.Items)
	if err != nil {
		s.updateOrderStatus(ctx, model.Order{ID: id, Status: model.StatusFailed})
		return 0, err
	}

	s.updateOrderStatus(ctx, model.Order{ID: id, Status: model.StatusAwaitingPayment})

	return id, nil
}

func (s *service) OrderGetInfo(ctx context.Context, orderID uint64) (model.Info, error) {
	order, err := s.repo.GetOrder(ctx, orderID)
	if err != nil {
		return model.Info{}, err
	}
	info := model.Info{
		UserID: order.UserID,
		Status: order.Status,
		Items:  order.Items,
	}

	return info, nil
}

func (s *service) OrderPay(ctx context.Context, orderID uint64) error {
	order, err := s.repo.GetOrder(ctx, orderID)
	if err != nil {
		return err
	}

	if order.Status != model.StatusAwaitingPayment {
		return errors.New(internalErrors.ErrOrderStatus)
	}

	err = s.stocksRemove(ctx, order.Items, false)
	if err != nil {
		return err
	}

	err = s.repo.UpdateOrderStatus(ctx, orderID, model.StatusPayed)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) OrderCancel(ctx context.Context, orderID uint64) error {
	order, err := s.repo.GetOrder(ctx, orderID)
	if err != nil {
		return err
	}

	if order.Status != model.StatusAwaitingPayment {
		return errors.New(internalErrors.ErrOrderStatus)
	}

	err = s.stocksRemove(ctx, order.Items, true)
	if err != nil {
		return err
	}

	err = s.repo.UpdateOrderStatus(ctx, orderID, model.StatusCanceled)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) OrderGetStockInfo(ctx context.Context, skuID uint32) (uint64, error) {
	stock := s.repo.GetStockBySKU(ctx, skuID)
	count := uint64(stock.TotalCount - stock.Reserved)

	return count, nil
}
