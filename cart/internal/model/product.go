package model

type Product struct {
	Token string `json:"token"`
	SKU   uint32 `json:"sku"`
}

type ProductRespV1 struct {
	Name    string `json:"name"`
	Price   uint32 `json:"price"`
	Code    uint32 `json:"code"`
	Message string `json:"message"`
	Details []any  `json:"details"`
}
