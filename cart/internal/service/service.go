package service

import (
	"context"
	"errors"
	"log"

	desc "route256.ozon.ru/project/cart/external/stocks/gen/api/orders/v1"
	internalErrors "route256.ozon.ru/project/cart/internal/errors"
	"route256.ozon.ru/project/cart/internal/model"
	"route256.ozon.ru/project/cart/util"
)

type Repository interface {
	CreateCart(ctx context.Context, userID int64) (model.Cart, error)
	GetCartByUserID(ctx context.Context, userID int64) (model.Cart, error)
	AddGoodToCart(ctx context.Context, userID int64, good model.Good) error
	RemoveGoodFromCart(ctx context.Context, userID int64, skuID uint32) error
	CleanUpCart(ctx context.Context, userID int64) error
}

type ProductClient interface {
	GetProduct(ctx context.Context, SKU uint32, count uint16) (*model.Good, error)
}

type StocksClient interface {
	GetStocksInfo(ctx context.Context, SKU uint32) (uint64, error)
	CreateOrder(ctx context.Context, order *desc.Order) (uint64, error)
}

type service struct {
	repo          Repository
	productClient ProductClient
	stocksClient  StocksClient
}

func New(repo Repository, productClient ProductClient, stocksClient StocksClient) *service {
	return &service{
		repo:          repo,
		productClient: productClient,
		stocksClient:  stocksClient,
	}
}

func (s *service) AddToCart(ctx context.Context, userID, skuID int64, count uint16) error {
	cart, err := s.repo.GetCartByUserID(ctx, userID)
	if err != nil {
		return err
	}
	if cart.ID <= 0 {
		cart, err = s.repo.CreateCart(ctx, userID)
		if err != nil {
			return err
		}
	}

	good, err := s.productClient.GetProduct(ctx, uint32(skuID), count)
	if err != nil {
		return err
	}
	if good != nil {
		good.Count = count

		lomsCount, err := s.stocksClient.GetStocksInfo(ctx, uint32(skuID))
		if err != nil {
			return err
		}
		if count > uint16(lomsCount) {
			return errors.New(internalErrors.ErrStockCount)
		}

		err = s.repo.AddGoodToCart(ctx, userID, *good)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *service) RemoveFromCart(ctx context.Context, userID, skuID int64) error {
	cart, err := s.repo.GetCartByUserID(ctx, userID)
	if err != nil {
		return err
	}
	if cart.ID <= 0 {
		return nil
	}

	sku := uint32(skuID)
	if sku == 0 {
		log.Println(internalErrors.ErrBadSKU)
		return nil
	}
	err = s.repo.RemoveGoodFromCart(ctx, userID, sku)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) CleanUpCart(ctx context.Context, userID int64) error {
	cart, err := s.repo.GetCartByUserID(ctx, userID)
	if err != nil {
		return err
	}
	if cart.ID <= 0 || len(cart.Goods) <= 0 {
		return nil
	}

	err = s.repo.CleanUpCart(ctx, userID)
	if err != nil {
		return err
	}

	return nil
}

func (s *service) GetCart(ctx context.Context, userID int64) (*model.Cart, error) {
	cart, err := s.repo.GetCartByUserID(ctx, userID)
	if err != nil {
		return nil, err
	}
	if cart.ID <= 0 || len(cart.Goods) <= 0 {
		log.Println(internalErrors.InfoFailedGetCart)
		return nil, nil
	}

	errGroups := make([]*util.ProductGroup, 0, len(cart.Goods))
	// errorGroup для каждого запроса c клиента GetProduct
	for _, g := range cart.Goods {
		eg, ctx := util.EGWithContext(ctx)
		eg.SetLimit(10)

		eg.GoPayload(func() (*model.Good, error) {
			defer func() {
				if err != nil {
					log.Println(err)
				}
			}()
			return s.productClient.GetProduct(ctx, g.SkuID, g.Count)
		})

		errGroups = append(errGroups, eg)
	}

	for _, eg := range errGroups {
		good, err := eg.Wait()

		if err != nil || good == nil {
			return nil, err
		} else {
			cart.TotalPrice = cart.TotalPrice + good.Price*uint32(good.Count)
		}
	}

	return &cart, nil
}

func (s *service) Checkout(ctx context.Context, userID int64) (model.Order, error) {
	cart, err := s.GetCart(ctx, userID)
	if err != nil {
		return model.Order{}, err
	}

	items := make([]*desc.Item, 0, len(cart.Goods))
	for _, e := range cart.Goods {
		items = append(items, &desc.Item{Sku: e.SkuID, Count: uint32(e.Count)})
	}

	orderID, err := s.stocksClient.CreateOrder(ctx, &desc.Order{
		UserId: userID,
		Items:  items,
	})
	if err != nil {
		return model.Order{}, err
	}

	err = s.CleanUpCart(ctx, userID)
	if err != nil {
		return model.Order{}, err
	}

	return model.Order{ID: orderID}, nil
}
