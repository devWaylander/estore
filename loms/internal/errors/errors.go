package errors

const (
	// ===================-   COMMON   -===================
	ErrMarshalResponse = "ERR_FAILED_TO_ENCODE_JSON_RESP"

	// ===================-    ORDER    -===================
	ErrCreateOrder   = "ERR_CREATE_FAIL"
	ErrOrderNotFound = "ERR_ORDER_NOT_FOUND"
	ErrOrderStatus   = "ERR_WRONG_ORDER_STATUS"

	// ===================-    STOCKS    -===================
	ErrStocksAreEmpty = "ERR_STOCKS_EMPTY"
	ErrStocksUpdate   = "ERR_STOCKS_UPDATE_FAILED"
)
