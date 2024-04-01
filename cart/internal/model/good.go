package model

type Good struct {
	SkuID uint32 `json:"sku_id"`

	Name  string `json:"name"`
	Price uint32 `json:"price"`
	Count uint16 `json:"count"`
}

type GoodRespV1 struct {
	SkuID int64 `json:"sku_id"`

	Name  string `json:"name"`
	Count uint16 `json:"count"`
	Price uint32 `json:"price"`
}

func (g *Good) ToGoodV1() *GoodRespV1 {
	return &GoodRespV1{
		SkuID: int64(g.SkuID),
		Name:  g.Name,
		Count: g.Count,
		Price: g.Price,
	}
}
