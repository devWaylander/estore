package service

import (
	"context"
	"errors"
	"testing"

	"github.com/gojuno/minimock/v3"
	"github.com/stretchr/testify/require"
	internalErrors "route256.ozon.ru/project/cart/internal/errors"
	"route256.ozon.ru/project/cart/internal/model"
)

const (
	skuID  = 773297411
	userID = 1
	cartID = userID + 1
)

type inputData struct {
	userID int64
	cartID int64
	skuID  int64
	count  uint16
	goods  map[uint32]model.Good
}

func TestAddToCart(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)
	productClientMock := NewProductClientMock(mc)
	stocksClientMock := NewStocksClientMock(mc)
	repoMock := NewRepositoryMock(mc)

	service := New(repoMock, productClientMock, stocksClientMock)

	tests := []struct {
		name      string
		inputData inputData
		wantErr   error
	}{
		{
			name: "success found product",
			inputData: inputData{
				userID: userID,
				skuID:  skuID,
				count:  10,
			},
			wantErr: nil,
		},
		{
			name: "failed found product",
			inputData: inputData{
				userID: userID,
				skuID:  123,
				count:  1,
			},
			wantErr: errors.New(internalErrors.ErrSKUNotFound),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			repoMock.GetCartByUserIDMock.Expect(ctx, tt.inputData.userID).Return(model.Cart{}, nil)
			repoMock.CreateCartMock.Expect(ctx, tt.inputData.userID).Return(model.Cart{}, nil)
			productClientMock.GetProductMock.Expect(ctx, uint32(tt.inputData.skuID)).Return(&model.Good{}, tt.wantErr)
			stocksClientMock.GetStocksInfoMock.Expect(ctx, uint32(tt.inputData.skuID)).Return(uint64(tt.inputData.count), tt.wantErr)
			if tt.wantErr == nil {
				repoMock.AddGoodToCartMock.Expect(ctx, tt.inputData.userID, model.Good{Count: tt.inputData.count}).Return(nil)
			}

			err := service.AddToCart(ctx, tt.inputData.userID, tt.inputData.skuID, tt.inputData.count)
			require.ErrorIs(t, err, tt.wantErr)
		})
	}
}

func TestRemoveFromCart(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)
	productClientMock := NewProductClientMock(mc)
	stocksClientMock := NewStocksClientMock(mc)
	repoMock := NewRepositoryMock(mc)

	service := New(repoMock, productClientMock, stocksClientMock)

	tests := []struct {
		name      string
		inputData inputData
		wantErr   error
	}{
		{
			name: "remove exists",
			inputData: inputData{
				userID: userID,
				cartID: cartID,
				skuID:  skuID,
			},
			wantErr: nil,
		},
		{
			name: "remove empty",
			inputData: inputData{
				userID: userID,
				cartID: 0,
				skuID:  0,
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			repoMock.GetCartByUserIDMock.Expect(ctx, tt.inputData.userID).Return(model.Cart{ID: tt.inputData.cartID}, nil)
			repoMock.RemoveGoodFromCartMock.Expect(ctx, tt.inputData.userID, uint32(tt.inputData.skuID)).Return(tt.wantErr)

			err := service.RemoveFromCart(ctx, tt.inputData.userID, tt.inputData.skuID)
			require.ErrorIs(t, err, tt.wantErr)
		})
	}
}

func TestCleanupCart(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)
	productClientMock := NewProductClientMock(mc)
	stocksClientMock := NewStocksClientMock(mc)
	repoMock := NewRepositoryMock(mc)

	service := New(repoMock, productClientMock, stocksClientMock)

	tests := []struct {
		name      string
		inputData inputData
		wantErr   error
	}{
		{
			name: "cleanup exists user cart",
			inputData: inputData{
				userID: userID,
				cartID: cartID,
				goods:  map[uint32]model.Good{1: {SkuID: 1}},
			},
			wantErr: nil,
		},
		{
			name: "cleanup empty user cart",
			inputData: inputData{
				userID: 0,
				cartID: 0,
				goods:  map[uint32]model.Good{},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			repoMock.GetCartByUserIDMock.Expect(ctx, tt.inputData.userID).Return(model.Cart{ID: tt.inputData.cartID, Goods: tt.inputData.goods}, nil)
			repoMock.CleanUpCartMock.Expect(ctx, tt.inputData.userID).Return(tt.wantErr)

			err := service.CleanUpCart(ctx, tt.inputData.userID)
			require.ErrorIs(t, err, tt.wantErr)
		})
	}
}

func TestGetCart(t *testing.T) {
	t.Parallel()

	mc := minimock.NewController(t)
	productClientMock := NewProductClientMock(mc)
	stocksClientMock := NewStocksClientMock(mc)
	repoMock := NewRepositoryMock(mc)

	service := New(repoMock, productClientMock, stocksClientMock)

	tests := []struct {
		name      string
		inputData inputData
		wantErr   error
	}{
		{
			name: "get exists user cart",
			inputData: inputData{
				userID: userID,
				cartID: cartID,
				goods: map[uint32]model.Good{
					1: {SkuID: 1, Name: "a", Price: 1, Count: 1},
					2: {SkuID: 2, Name: "b", Price: 2, Count: 2},
				},
			},
			wantErr: nil,
		},
		{
			name: "get empty user cart",
			inputData: inputData{
				userID: 0,
				cartID: 0,
				goods:  map[uint32]model.Good{},
			},
			wantErr: errors.New(internalErrors.InfoFailedGetCart),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()

			repoMock.GetCartByUserIDMock.Expect(ctx, tt.inputData.userID).Return(model.Cart{ID: tt.inputData.cartID, Goods: tt.inputData.goods}, tt.wantErr)
			for i := range tt.inputData.goods {
				productClientMock.GetProductMock.When(ctx, tt.inputData.goods[i].SkuID).Then(&model.Good{}, tt.wantErr)
			}

			_, err := service.GetCart(ctx, tt.inputData.userID)
			require.ErrorIs(t, err, tt.wantErr)
		})
	}
}
