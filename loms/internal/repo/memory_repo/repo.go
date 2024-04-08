package memory_repo

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"math/rand"
	"os"
	"sync"

	internalErrors "route256.ozon.ru/project/loms/internal/errors"
	"route256.ozon.ru/project/loms/internal/model"
)

type repository struct {
	ordersStorage map[uint64]model.Order
	ordersLock    sync.RWMutex

	stocksStorage map[uint32]model.Stock
	stocksLock    sync.RWMutex
}

func New(jsonStock *os.File) *repository {
	repo := repository{
		ordersStorage: map[uint64]model.Order{},
		stocksStorage: map[uint32]model.Stock{},
	}

	repo.initStock(jsonStock)
	return &repo
}

// TODO: временное решение, пока СУБД не выдаёт ID
func randomUint64() uint64 {
	return uint64(rand.Uint32())<<32 + uint64(rand.Uint32())
}

func (r *repository) initStock(jsonStock *os.File) {
	byteVal, err := io.ReadAll(jsonStock)
	if err != nil {
		log.Panic(err)
	}

	tmp := []model.Stock{}
	json.Unmarshal(byteVal, &tmp)

	r.stocksLock.Lock()
	defer r.stocksLock.Unlock()
	for _, e := range tmp {
		r.stocksStorage[e.SKU] = e
	}
}

func (r *repository) CreateOrder(ctx context.Context, data model.Order) (uint64, error) {
	data.ID = randomUint64()

	r.ordersLock.Lock()
	defer r.ordersLock.Unlock()

	r.ordersStorage[data.ID] = data

	if r.ordersStorage[data.ID].ID <= 0 {
		return 0, errors.New(internalErrors.ErrCreateOrder)
	}

	return data.ID, nil
}

func (r *repository) GetOrder(ctx context.Context, orderID uint64) (model.Order, error) {
	r.ordersLock.RLock()
	defer r.ordersLock.RUnlock()

	order := r.ordersStorage[orderID]
	if order.ID == 0 {
		return model.Order{}, errors.New(internalErrors.ErrOrderNotFound)
	}

	return order, nil
}

func (r *repository) UpdateOrderStatus(ctx context.Context, orderID uint64, status string) error {
	r.ordersLock.Lock()
	defer r.ordersLock.Unlock()

	tmp := r.ordersStorage[orderID]
	tmp.Status = status
	r.ordersStorage[orderID] = tmp

	return nil
}

func (r *repository) GetStockBySKU(ctx context.Context, SKU uint32) model.Stock {
	r.stocksLock.RLock()
	defer r.stocksLock.RUnlock()

	return r.stocksStorage[SKU]
}

func (r *repository) UpdateStock(ctx context.Context, data model.Stock) error {
	r.stocksLock.Lock()
	defer r.stocksLock.Unlock()

	r.stocksStorage[data.SKU] = data
	if r.stocksStorage[data.SKU].TotalCount != data.TotalCount ||
		r.stocksStorage[data.SKU].Reserved != data.Reserved {
		return errors.New(internalErrors.ErrStocksUpdate)
	}

	return nil
}
