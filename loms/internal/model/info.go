package model

import servicepb "route256.ozon.ru/project/loms/gen/api/orders/v1"

type Info struct {
	Status string `json:"status"`
	UserID int64  `json:"user_id"`
	Items  []Item `json:"items"`
}

func ToOrderGetInfoRespServicepb(info Info) *servicepb.OrderGetInfoResp {
	protoInfoResp := servicepb.OrderGetInfoResp{
		UserId: info.UserID,
		Status: servicepb.Statuses(StatusesValue[info.Status]),
	}

	if info.Items != nil {
		protoInfoResp.Items = make([]*servicepb.Item, len(info.Items))
		for i, e := range info.Items {
			// TODO: protobuf имеет только uint32
			// https://protobuf.dev/programming-guides/proto3/
			protoInfoResp.Items[i] = &servicepb.Item{
				Count: uint32(e.Count),
				Sku:   e.SKU,
			}
		}
	}

	return &protoInfoResp
}
