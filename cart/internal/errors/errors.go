package internalErrors

const (
	// ===================-   COMMON   -===================
	ErrMarshalResponse = "ERR_FAILED_TO_ENCODE_JSON_RESP"

	// ===================-    USER    -===================
	ErrMissingUserID = "ERR_MISSING_USER_ID"

	// ===================-    FILES   -===================
	ErrOpenFile        = "ERR_FAILED_TO_OPEN_FILE"
	ErrWriteFile       = "ERR_FAILED_TO_WRITE_FILE"
	ErrCreateDirectory = "ERR_FAILED_TO_CREATE_DIRECTORY"
	ErrFileStat        = "ERR_FAILED_TO_GET_FILE_STAT"
	ErrConvertFile     = "ERR_FAILED_TO_CONVERT_FILE"

	// ===================-    CART    -===================
	ErrCartNotFound             = "ERR_CART_NOT_FOUND"
	ErrCountIsMissingOrNotValid = "ERR_COUNT_IS_MISSING_OR_NOT_VALID"
	ErrFailedAddAGood           = "ERR_FAILED_POST_A_GOOD"
	ErrFailedDeleteAGood        = "ERR_FAILED_DELETE_A_GOOD"
	ErrFailedCleanUpACart       = "ERR_FAILED_DELETE_A_GOODS"
	ErrFailedGetCart            = "ERR_FAILED_GET_CART"
	InfoFailedGetCart           = "INFO_CART_DOESNT_EXIST_OR_EMPTY"

	// ===================-  PRODUCT  -===================
	ErrMissingSKU    = "ERR_MISSING_SKU"
	ErrTokenNotValid = "ERR_TOKEN_NOT_VALID"
	ErrSKUNotFound   = "ERR_SKU_NOT_FOUND"
	ErrBadSKU        = "ERR_BAD_SKU"

	// ===================-    LOMS    -===================
	ErrStock      = "ERR_STOCK_FAILED"
	ErrStockCount = "ERR_STOCK_COUNT_FAILED_PRECONDITION"
	ErrOrder      = "ERR_ORDER_FAILED"
)
