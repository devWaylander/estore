package model

import servicepb "route256.ozon.ru/project/loms/gen/api/orders/v1"

const (
	StatusNew             = "STATUS_NEW"
	StatusAwaitingPayment = "STATUS_AWAITING_PAYMENT"
	StatusFailed          = "STATUS_FAILED"
	StatusPayed           = "STATUS_PAYED"
	StatusCanceled        = "STATUS_CANCELED"
)

var (
	StatusesValue = map[string]int32{
		"STATUS_UNSPECIFIED":  0,
		StatusNew:             1,
		StatusAwaitingPayment: 2,
		StatusFailed:          3,
		StatusPayed:           4,
		StatusCanceled:        5,
	}
)

type Order struct {
	ID     uint64 `json:"id"`
	UserID int64  `json:"user_id"`
	Items  []Item `json:"items"`
	Status string `json:"status"`
}

func ToOrder(protoOrder *servicepb.Order) Order {
	order := Order{
		UserID: protoOrder.UserId,
	}

	if protoOrder.Items != nil {
		order.Items = make([]Item, len(protoOrder.Items))
		for i, e := range protoOrder.Items {
			// TODO: protobuf имеет только uint32
			// https://protobuf.dev/programming-guides/proto3/
			order.Items[i].Count = uint16(e.Count)
			order.Items[i].SKU = e.Sku
		}
	}

	return order
}
