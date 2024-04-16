package gotripay

type ITripayContracts interface {
	Instruction(req *TripayIntructionParamsRequest) (*TripayInstructionResponse, error)

	// Open Payment
	RequestTransaction(req *TripayOpenPaymentCreateRequest) (*TripayOpenPaymentCreateResponse, error)
	DetailTransaction()
	ListTransaction()
}
