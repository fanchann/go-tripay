package gotripay

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type clientRegister struct {
	Client *Client
}

func NewClientRegister(client *Client) ITripayContracts {
	return &clientRegister{client}
}

func (c *clientRegister) Instruction(req *TripayIntructionParamsRequest) (*TripayInstructionResponse, error) {
	params := url.Values{}
	params.Set("code", string(req.BANK_CODE_CHANNEL))

	if req.PAYCODE != "" {
		params.Set("pay_code", req.PAYCODE)
	}

	if req.AMOUNT != "" {
		params.Set("amount", req.AMOUNT)
	}

	if req.ALLOW_HTML != "" {
		params.Set("allow_html", req.ALLOW_HTML)
	}

	queryStr := params.Encode()

	request := tripayRequestParam{
		URL:    URL(c.Client.URL()) + "payment/instruction?" + URL(queryStr),
		METHOD: HTTP_METHOD(http.MethodGet),
		BODY:   nil,
		HEADER: c.Client.HeaderRequest(),
	}

	response, err := request.Do()
	if err != nil {
		return nil, err
	}
	var responseBody map[string]interface{}
	if err := json.Unmarshal(response.ResponseBody, &responseBody); err != nil {
		return nil, err

	}

	data := responseBody["data"].([]interface{})

	return &TripayInstructionResponse{
		Success: responseBody["success"].(bool),
		Message: responseBody["message"].(string),
		Steps:   data,
	}, nil
}

// Open Payment
func (c *clientRegister) RequestTransaction(req *TripayOpenPaymentCreateRequest) (*TripayOpenPaymentCreateResponse, error) {
	payload := map[string]interface{}{
		"method":        req.METHOD + "OP",
		"merchant_ref":  req.MERCHANT_REF,
		"customer_name": req.CUSTOMER_NAME,
		"signature":     req.SIGNATURE,
	}

	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	request := tripayRequestParam{
		URL:    URL(c.Client.URL()) + "open-payment/create",
		METHOD: http.MethodPost,
		BODY:   payloadBytes,
		HEADER: c.Client.HeaderRequest(),
	}
	fmt.Printf("request.URL: %v\n", request.URL)

	response, err := request.Do()
	if err != nil {
		return nil, err
	}
	fmt.Printf("response: %v\n", string(response.ResponseBody))

	var responseBody map[string]interface{}
	if err := json.Unmarshal(response.ResponseBody, &responseBody); err != nil {
		return nil, err

	}

	return &TripayOpenPaymentCreateResponse{
		Success: responseBody["success"].(bool),
		Message: responseBody["message"].(string),
	}, nil
}

func (c *clientRegister) DetailTransaction() {
	panic("not implemented") // TODO: Implement
}

func (c *clientRegister) ListTransaction() {
	panic("not implemented") // TODO: Implement
}
