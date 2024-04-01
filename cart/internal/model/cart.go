package model

import "sort"

type Cart struct {
	ID     int64 `json:"id"`
	UserID int64 `json:"user_id"`

	Goods map[uint32]Good `json:"goods"`

	TotalPrice uint32 `json:"total_price"`
}

type CartRespV1 struct {
	Items      []*GoodRespV1 `json:"items"`
	TotalPrice uint32        `json:"total_price"`
}

func (c *Cart) ToCartRespV1() *CartRespV1 {
	cartRespV1 := CartRespV1{
		TotalPrice: c.TotalPrice,
		Items:      make([]*GoodRespV1, 0, len(c.Goods)),
	}

	for _, g := range c.Goods {
		cartRespV1.Items = append(cartRespV1.Items, g.ToGoodV1())
	}

	sort.Slice(cartRespV1.Items, func(i, j int) bool {
		return cartRespV1.Items[i].SkuID < cartRespV1.Items[j].SkuID
	})

	return &cartRespV1
}
