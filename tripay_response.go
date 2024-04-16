package gotripay

type TripayInstructionResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Steps   any    `json:"steps"`
}

type TripayOpenPaymentCreateResponse struct {
	Success bool                   `json:"success"`
	Message string                 `json:"message"`
	Data    map[string]interface{} `json:"data"`
}

type TripayResponse struct {
	HttpCode     int
	ResponseBody []byte
}
