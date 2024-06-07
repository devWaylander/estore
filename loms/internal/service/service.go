package service

import (
	"context"
	"errors"
	"log"

	"github.com/IBM/sarama"
	"github.com/jackc/pgx/v5/pgtype"
	internalErrors "route256.ozon.ru/project/loms/internal/errors"
	"route256.ozon.ru/project/loms/internal/infra/kafka"
	"route256.ozon.ru/project/loms/internal/infra/kafka/events"
	"route256.ozon.ru/project/loms/internal/infra/kafka/messages"
	"route256.ozon.ru/project/loms/internal/model"
	orders_repo "route256.ozon.ru/project/loms/internal/repo/db_repo/orders"
	stocks_repo "route256.ozon.ru/project/loms/internal/repo/db_repo/stocks"
)

type MemoryRepository interface {
	CreateOrder(ctx context.Context, data model.Order) (uint64, error)
	GetOrder(ctx context.Context, orderID uint64) (model.Order, error)
	UpdateOrderStatus(ctx context.Context, orderID uint64, status string) error
	GetStockBySKU(ctx context.Context, SKU uint32) model.Stock
	UpdateStock(ctx context.Context, data model.Stock) error
}

type OrderRepository interface {
	CreateItem(ctx context.Context, arg orders_repo.CreateItemParams) error
	GetItemsByOrderID(ctx context.Context, orderID int64) ([]orders_repo.GetItemsByOrderIDRow, error)
	CreateOrder(ctx context.Context, userID int64) (int64, error)
	GetOrder(ctx context.Context, id int64) (orders_repo.GetOrderRow, error)
	UpdateOrder(ctx context.Context, arg orders_repo.UpdateOrderParams) error
}

type StockRepository interface {
	CreateStock(ctx context.Context, arg stocks_repo.CreateStockParams) error
	GetStockBySKU(ctx context.Context, sku int32) (stocks_repo.StocksStock, error)
	UpdateStock(ctx context.Context, arg stocks_repo.UpdateStockParams) error
}

type service struct {
	ordersRepo   OrderRepository
	stocksRepo   StockRepository
	syncProducer sarama.SyncProducer
	kafkaConfig  kafka.Config
}

func New(ordersRepo OrderRepository, stocksRepo StockRepository, syncProducer sarama.SyncProducer, config kafka.Config) *service {
	return &service{
		ordersRepo:   ordersRepo,
		stocksRepo:   stocksRepo,
		syncProducer: syncProducer,
		kafkaConfig:  config,
	}
}

func (s *service) getItemsByOrderID(ctx context.Context, orderID uint64) ([]model.Item, error) {
	dbItems, err := s.ordersRepo.GetItemsByOrderID(ctx, int64(orderID))
	if err != nil {
		return []model.Item{}, err
	}

	items := make([]model.Item, len(dbItems))
	for i, e := range dbItems {
		items[i] = model.Item{
			SKU:   uint32(e.Sku),
			Count: uint16(e.Count),
		}
	}

	return items, nil
}

func (s *service) updateOrderStatus(ctx context.Context, orderID uint64, status string) error {
	err := s.ordersRepo.UpdateOrder(ctx, orders_repo.UpdateOrderParams{
		ID:     int64(orderID),
		Status: pgtype.Text{String: status, Valid: true},
	})
	if err != nil {
		log.Println(err)
		return err
	}

	eventFactory := events.NewDefaultFactory(int(orderID))
	event := eventFactory.Create(status)

	messageFactory := messages.NewDefaultFactory()
	message := messageFactory.Create(event, s.kafkaConfig)

	partition, offset, err := s.syncProducer.SendMessage(message)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("key: %d, partition: %d, offset: %d", event.ID, partition, offset)

	return nil
}

func (s *service) stocksReserve(ctx context.Context, data []model.Item) error {
	for _, e := range data {
		count := int32(e.Count)
		stock, err := s.stocksRepo.GetStockBySKU(ctx, int32(e.SKU))
		if err != nil {
			return err
		}
		if count > (stock.TotalCount - stock.Reserved) {
			return errors.New(internalErrors.ErrStocksAreEmpty)
		}

		stock.Reserved = stock.Reserved + count
		err = s.stocksRepo.UpdateStock(ctx, stocks_repo.UpdateStockParams{
			ID:         stock.ID,
			Sku:        stock.Sku,
			Reserved:   stock.Reserved,
			TotalCount: stock.TotalCount,
		})
		if err != nil {
			log.Println(err)
			return err
		}
	}

	return nil
}

func (s *service) stocksRemove(ctx context.Context, data []model.Item, isCancel bool) error {
	for _, e := range data {
		count := int32(e.Count)
		stock, err := s.stocksRepo.GetStockBySKU(ctx, int32(e.SKU))
		if err != nil {
			return err
		}

		stock.Reserved = stock.Reserved - count

		if !isCancel {
			stock.TotalCount = stock.TotalCount - count
		}

		err = s.stocksRepo.UpdateStock(ctx, stocks_repo.UpdateStockParams{
			ID:         stock.ID,
			Sku:        stock.Sku,
			Reserved:   stock.Reserved,
			TotalCount: stock.TotalCount,
		})
		if err != nil {
			log.Println(err)
			return err
		}
	}

	return nil
}

func (s *service) OrderCreate(ctx context.Context, data model.Order) (uint64, error) {
	data.Status = model.StatusNew
	id, err := s.ordersRepo.CreateOrder(ctx, data.UserID)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	err = s.stocksReserve(ctx, data.Items)
	if err != nil {
		s.updateOrderStatus(ctx, uint64(id), model.StatusFailed)
		return 0, err
	}

	err = s.updateOrderStatus(ctx, uint64(id), model.StatusAwaitingPayment)
	if err != nil {
		return 0, err
	}

	for _, e := range data.Items {
		err := s.ordersRepo.CreateItem(ctx, orders_repo.CreateItemParams{
			OrderID: int64(id),
			Sku:     int32(e.SKU),
			Count:   int32(e.Count),
		})
		if err != nil {
			return uint64(0), err
		}
	}

	return uint64(id), nil
}

func (s *service) OrderGetInfo(ctx context.Context, orderID uint64) (model.Info, error) {
	order, err := s.ordersRepo.GetOrder(ctx, int64(orderID))
	if err != nil {
		return model.Info{}, err
	}

	items, err := s.getItemsByOrderID(ctx, orderID)
	if err != nil {
		return model.Info{}, err
	}

	info := model.Info{
		UserID: order.UserID,
		Status: order.Status.String,
		Items:  items,
	}

	return info, nil
}

func (s *service) OrderPay(ctx context.Context, orderID uint64) error {
	order, err := s.ordersRepo.GetOrder(ctx, int64(orderID))
	if err != nil {
		return err
	}

	if order.Status.String != model.StatusAwaitingPayment {
		return errors.New(internalErrors.ErrOrderStatus)
	}

	items, err := s.getItemsByOrderID(ctx, orderID)
	if err != nil {
		return err
	}

	err = s.stocksRemove(ctx, items, false)
	if err != nil {
		return err
	}

	err = s.updateOrderStatus(ctx, orderID, model.StatusPayed)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) OrderCancel(ctx context.Context, orderID uint64) error {
	order, err := s.ordersRepo.GetOrder(ctx, int64(orderID))
	if err != nil {
		return err
	}

	if order.Status.String != model.StatusAwaitingPayment {
		return errors.New(internalErrors.ErrOrderStatus)
	}

	items, err := s.getItemsByOrderID(ctx, orderID)
	if err != nil {
		return err
	}

	err = s.stocksRemove(ctx, items, true)
	if err != nil {
		return err
	}

	err = s.updateOrderStatus(ctx, orderID, model.StatusCanceled)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) OrderGetStockInfo(ctx context.Context, skuID uint32) (uint64, error) {
	stock, err := s.stocksRepo.GetStockBySKU(ctx, int32(skuID))
	if err != nil {
		return 0, err
	}

	count := uint64(stock.TotalCount - stock.Reserved)

	return count, nil
}
