package gotripay

type Mode string
type URL string

var DEVELOPMENT Mode = "development"
var PRODUCTION Mode = "production"

const URL_DEVELOPMENT URL = "https://tripay.co.id/api-sandbox/"
const URL_PRODUCTION URL = "https://tripay.co.id/api/"

type Client struct {
	MerchantCode string
	ApiKey       string
	PrivateKey   string
	Mode         Mode
}

func (c *Client) HeaderRequest() []map[string]string {
	headers := []map[string]string{
		{"Content-Type": "application/json"},
		{"Authorization": "Bearer " + c.ApiKey},
	}
	return headers
}

func (c *Client) URL() string {
	if c.Mode == DEVELOPMENT {
		return string(URL_DEVELOPMENT)
	}
	return string(URL_PRODUCTION)
}
