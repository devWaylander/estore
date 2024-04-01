package repo

import (
	"context"
	"sync"

	"route256.ozon.ru/project/cart/internal/model"
)

type repository struct {
	items     map[int64]model.Cart
	itemsLock sync.RWMutex
}

func New() *repository {
	return &repository{
		items: map[int64]model.Cart{},
	}
}

func (r *repository) CreateCart(ctx context.Context, userID int64) (model.Cart, error) {
	r.itemsLock.Lock()
	defer r.itemsLock.Unlock()

	cart := model.Cart{
		// temporary solution, while we have 1 user -> 1 cart
		ID:         userID + 1,
		UserID:     userID,
		Goods:      map[uint32]model.Good{},
		TotalPrice: uint32(0),
	}

	r.items[cart.UserID] = cart

	return cart, nil
}

func (r *repository) GetCartByUserID(ctx context.Context, userID int64) (model.Cart, error) {
	r.itemsLock.RLock()
	defer r.itemsLock.RUnlock()

	if cart, found := r.items[userID]; found {
		return cart, nil
	}
	return model.Cart{}, nil
}

func (r *repository) AddGoodToCart(ctx context.Context, userID int64, good model.Good) error {
	r.itemsLock.Lock()
	defer r.itemsLock.Unlock()
	if good.SkuID <= 0 {
		return nil
	}

	if r.items[userID].Goods[good.SkuID].SkuID == good.SkuID {
		entry := r.items[userID].Goods[good.SkuID]
		entry.Count = entry.Count + good.Count
		r.items[userID].Goods[good.SkuID] = entry
	} else {
		r.items[userID].Goods[good.SkuID] = good
	}

	return nil
}

func (r *repository) RemoveGoodFromCart(ctx context.Context, userID int64, skuID uint32) error {
	r.itemsLock.Lock()
	defer r.itemsLock.Unlock()

	if r.items[userID].Goods == nil {
		return nil
	}

	if r.items[userID].Goods[skuID].SkuID == skuID {
		delete(r.items[userID].Goods, skuID)
	} else {
		return nil
	}

	return nil
}

func (r *repository) CleanUpCart(ctx context.Context, userID int64) error {
	r.itemsLock.Lock()
	defer r.itemsLock.Unlock()

	entry := r.items[userID]
	entry.Goods = map[uint32]model.Good{}
	r.items[userID] = entry

	return nil
}
