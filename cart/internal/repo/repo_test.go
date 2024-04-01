package repo

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
	"route256.ozon.ru/project/cart/internal/model"
)

const (
	// magic 27 for cpu=16 :)
	parallelism = 27
	skuID       = 31147466
	userID      = 1
	cartID      = userID + 1
)

type inputData struct {
	userID int64
	skuID  int64
	count  uint16
	cart   model.Cart
	good   model.Good
}

func BenchmarkAddGoodToCart(b *testing.B) {
	ctx := context.Background()

	inputData := inputData{
		userID: userID,
		skuID:  skuID,
		count:  10,
		good: model.Good{
			SkuID: skuID,
			Name:  "Good good",
			Price: 123,
			Count: 1,
		},
	}

	b.SetParallelism(parallelism)
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			repo := New()
			repo.CreateCart(ctx, inputData.userID)

			repo.AddGoodToCart(ctx, inputData.userID, inputData.good)
		}
	})
}

func TestCreateCart(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		inputData inputData
		wantErr   error
	}{
		{
			name: "success create cart",
			inputData: inputData{
				userID: userID,
				skuID:  skuID,
				count:  10,
				cart: model.Cart{
					ID:         cartID,
					UserID:     userID,
					Goods:      map[uint32]model.Good{},
					TotalPrice: uint32(0),
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()

			repo := New()

			cart, _ := repo.CreateCart(ctx, tt.inputData.userID)
			require.EqualValues(t, tt.inputData.cart, cart)
		})
	}
}

func TestGetCartByUserID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		inputData inputData
		wantErr   error
	}{
		{
			name: "success get cart",
			inputData: inputData{
				userID: userID,
				skuID:  skuID,
				count:  10,
				cart: model.Cart{
					ID:         cartID,
					UserID:     userID,
					Goods:      map[uint32]model.Good{},
					TotalPrice: uint32(0),
				},
			},
			wantErr: nil,
		},
		{
			name: "failed get cart",
			inputData: inputData{
				userID: 0,
				skuID:  skuID,
				count:  10,
				cart: model.Cart{
					ID:         cartID,
					UserID:     userID,
					Goods:      map[uint32]model.Good{},
					TotalPrice: uint32(0),
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()

			repo := New()

			cart, _ := repo.CreateCart(ctx, tt.inputData.userID)
			cart, _ = repo.GetCartByUserID(ctx, tt.inputData.userID)
			require.EqualValues(t, model.Cart{
				ID:         tt.inputData.userID + 1,
				UserID:     tt.inputData.userID,
				Goods:      map[uint32]model.Good{},
				TotalPrice: uint32(0),
			}, cart)
		})
	}
}

func TestAddGoodToCart(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		inputData inputData
		wantErr   error
	}{
		{
			name: "success add good to cart",
			inputData: inputData{
				userID: userID,
				skuID:  skuID,
				count:  10,
				cart: model.Cart{
					ID:     cartID,
					UserID: userID,
					Goods: map[uint32]model.Good{skuID: {
						SkuID: skuID,
						Name:  "Good good",
						Price: 123,
						Count: 1,
					}},
					TotalPrice: uint32(0),
				},
				good: model.Good{
					SkuID: skuID,
					Name:  "Good good",
					Price: 123,
					Count: 1,
				},
			},
			wantErr: nil,
		},
		{
			name: "failed add good to cart",
			inputData: inputData{
				userID: 0,
				skuID:  skuID,
				count:  10,
				cart: model.Cart{
					ID:         1,
					UserID:     0,
					Goods:      map[uint32]model.Good{},
					TotalPrice: uint32(0),
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()

			repo := New()

			cart, _ := repo.CreateCart(ctx, tt.inputData.userID)
			err := repo.AddGoodToCart(ctx, tt.inputData.userID, tt.inputData.good)
			cart, _ = repo.GetCartByUserID(ctx, tt.inputData.userID)

			require.EqualValues(t, tt.inputData.cart, cart)
			require.ErrorIs(t, err, tt.wantErr)
		})
	}
}

func TestRemoveGoodFromCart(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		inputData inputData
		wantErr   error
	}{
		{
			name: "success remove good from cart",
			inputData: inputData{
				userID: 1,
				skuID:  skuID,
				count:  10,
				cart: model.Cart{
					ID:         cartID,
					UserID:     1,
					Goods:      map[uint32]model.Good{},
					TotalPrice: uint32(0),
				},
				good: model.Good{
					SkuID: skuID,
					Name:  "Good good",
					Price: 123,
					Count: 1,
				},
			},
			wantErr: nil,
		},
		{
			name: "failed remove good from cart",
			inputData: inputData{
				userID: 0,
				skuID:  0,
				count:  10,
				cart: model.Cart{
					ID:     1,
					UserID: 0,
					Goods: map[uint32]model.Good{skuID: {
						SkuID: skuID,
						Name:  "Good good",
						Price: 123,
						Count: 1,
					}}, TotalPrice: uint32(0),
				},
				good: model.Good{
					SkuID: skuID,
					Name:  "Good good",
					Price: 123,
					Count: 1,
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()

			repo := New()

			cart, _ := repo.CreateCart(ctx, tt.inputData.userID)
			err := repo.AddGoodToCart(ctx, tt.inputData.userID, tt.inputData.good)
			err = repo.RemoveGoodFromCart(ctx, tt.inputData.userID, uint32(tt.inputData.skuID))
			cart, _ = repo.GetCartByUserID(ctx, tt.inputData.userID)

			require.EqualValues(t, tt.inputData.cart, cart)
			require.ErrorIs(t, err, tt.wantErr)
		})
	}
}

func TestCleanUpCart(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		inputData inputData
		wantErr   error
	}{
		{
			name: "success remove good from cart",
			inputData: inputData{
				userID: 1,
				skuID:  skuID,
				count:  10,
				cart: model.Cart{
					ID:         cartID,
					UserID:     1,
					Goods:      map[uint32]model.Good{},
					TotalPrice: uint32(0),
				},
				good: model.Good{
					SkuID: skuID,
					Name:  "Good good",
					Price: 123,
					Count: 1,
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctx := context.Background()

			repo := New()

			cart, _ := repo.CreateCart(ctx, tt.inputData.userID)
			err := repo.AddGoodToCart(ctx, tt.inputData.userID, tt.inputData.good)
			err = repo.AddGoodToCart(ctx, tt.inputData.userID, tt.inputData.good)
			err = repo.CleanUpCart(ctx, tt.inputData.userID)
			cart, _ = repo.GetCartByUserID(ctx, tt.inputData.userID)

			require.EqualValues(t, tt.inputData.cart, cart)
			require.ErrorIs(t, err, tt.wantErr)
		})
	}
}
