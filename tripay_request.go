package gotripay

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

type HTTP_METHOD string

type tripayRequestParam struct {
	URL    URL
	METHOD HTTP_METHOD
	BODY   []byte
	HEADER []map[string]string
}

func (t *tripayRequestParam) Do() (*TripayResponse, error) {
	req, err := http.NewRequest(string(t.METHOD), string(t.URL), bytes.NewBuffer(t.BODY))
	if err != nil {
		return nil, err
	}

	for _, header := range t.HEADER {
		for key, value := range header {
			req.Header.Set(key, value)
		}
	}

	// Create a new HTTP client
	client := &http.Client{}

	// Send the request
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Read the response body
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &TripayResponse{
		HttpCode:     resp.StatusCode,
		ResponseBody: responseBody,
	}, nil
}
