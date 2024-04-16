package gotripay

type TripayIntructionParamsRequest struct {
	BANK_CODE_CHANNEL TRIPAY_CHANNEL
	PAYCODE           string
	AMOUNT            string
	ALLOW_HTML        string
}

type TripayOpenPaymentCreateRequest struct {
	METHOD        TRIPAY_CHANNEL
	MERCHANT_REF  string
	CUSTOMER_NAME string
	SIGNATURE     TRIPAY_SIGNATURE
}
