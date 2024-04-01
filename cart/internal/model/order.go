package model

type Order struct {
	ID uint64 `json:"id"`
}

type OrderRespV1 struct {
	ID uint64 `json:"order_id"`
}

func (o *Order) ToOrderV1() *OrderRespV1 {
	return &OrderRespV1{
		ID: o.ID,
	}
}
