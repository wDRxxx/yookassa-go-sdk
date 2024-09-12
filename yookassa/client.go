package yookassa

import (
	"bytes"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

const (
	baseURL = "https://api.yookassa.ru/v3"
)

type Client struct {
	shopID  string
	shopKey string
	client  *http.Client
}

func NewClient(shopID string, shopKey string) *Client {
	return &Client{
		shopID:  shopID,
		shopKey: shopKey,
		client:  &http.Client{},
	}
}

func (c *Client) sendRequest(endpoint string, method string, body []byte) (*http.Response, error) {
	url := fmt.Sprintf("%s%s", baseURL, endpoint)
	req, err := http.NewRequest(method, url, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.SetBasicAuth(c.shopID, c.shopKey)

	if method == http.MethodPost {
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("Idempotence-Key", uuid.NewString())
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
