package main

import (
	"fmt"

	gotripay "github.com/zakirkun/go-tripay"
)

func main() {
	// tripayInstruction()
	createPayment()
}

func tripayInstruction() {
	client := gotripay.Client{
		MerchantCode: "T30863",
		ApiKey:       "DEV-LUQxEmU4JNlsX4jCkPGhoFLX8DoGyQTg6638b0Ad",
		PrivateKey:   "UBDvO-ISKtj-pOtTp-Cv7XU-tycTT",
		Mode:         gotripay.DEVELOPMENT,
	}

	clientService := gotripay.NewClientRegister(&client)

	instructionParam := gotripay.TripayIntructionParamsRequest{
		BANK_CODE_CHANNEL: gotripay.CHANNEL_BRIVA,
		PAYCODE:           "",
		AMOUNT:            "100000",
		ALLOW_HTML:        "",
	}
	instructionResponse, err := clientService.Instruction(&instructionParam)
	if err != nil {
		panic(err)
	}

	fmt.Printf("instructionResponse: %v\n", instructionResponse)
}

func createPayment() {
	client := gotripay.Client{
		MerchantCode: "T14302",
		ApiKey:       "DEV-ZKIDl5gE3AsCDThj7mWX6yvQ8f42NZWJWlZ7TSzS",
		PrivateKey:   "J2WTm-93avv-w0PZV-ur1t4-4TCjd",
		Mode:         gotripay.DEVELOPMENT,
	}

	clientService := gotripay.NewClientRegister(&client)

	// create signature
	signatureStr := gotripay.Signature{
		Amount:       100000,
		PrivateKey:   client.PrivateKey,
		MerchantRef:  "INV345678",
		MerchantCode: "T0001",
		Channel:      gotripay.CHANNEL_BRIVA,
	}

	instructionParam := gotripay.TripayOpenPaymentCreateRequest{
		METHOD:        gotripay.CHANNEL_BRIVA,
		MERCHANT_REF:  "",
		CUSTOMER_NAME: "FardaAyu",
		SIGNATURE:     signatureStr.CreateSignature(),
	}
	instructionResponse, err := clientService.RequestTransaction(&instructionParam)
	if err != nil {
		panic(err)
	}

	fmt.Printf("instructionResponse: %v\n", instructionResponse)
}
