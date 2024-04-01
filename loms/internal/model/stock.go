package model

type Stock struct {
	SKU        uint32 `json:"sku"`
	TotalCount uint16 `json:"total_count"`
	Reserved   uint16 `json:"reserved"`
}
