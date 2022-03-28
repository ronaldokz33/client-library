package interview

import (
	"net/http"
	"time"
	"fmt"
	"errors"
	"encoding/json"
)

const baseURLV1 = "http://localhost:8080/v1"

//Client is a helper to handle with the api calls
type Client struct {
	baseURL string
	HTTPClient *http.Client
}

//NewClient is used to initialize the client that will be used to call the API.
func NewClient() *Client  {
	return &Client {
		baseURL: baseURLV1,
		HTTPClient: &http.Client {
			Timeout: time.Minute,
		},
	}
}

//successResponse is a helper to handle with success responses
type successResponse struct {
	Code int `json:"code"`
	Data interface{} `json:"data"`
}

//errorReponse is a helper to handle with error responses
type errorReponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
}

//sendRequest is responsable to handle with the API request
//if you need to create some authorization you do here
func (c *Client) sendRequest (req *http.Request, v interface{}) error {
	req.Header.Set("Accept", "application/json; charset=utf-8")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	res, err := c.HTTPClient.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()
	if res.StatusCode < 200 || res.StatusCode >= 300  {
		var errRes errorReponse
		if  err := json.NewDecoder(res.Body).Decode(&errRes); err == nil {
			return errors.New(errRes.Message)
		}

		return fmt.Errorf("unknown error, status code: %d", res.StatusCode)
	}

	fullResponse := successResponse {
		Data: v,
	}

	if res.StatusCode != http.StatusNoContent {
		if err = json.NewDecoder(res.Body).Decode(&fullResponse); err != nil {
			return err
		}
	}

	return nil
}