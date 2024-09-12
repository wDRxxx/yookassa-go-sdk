package yookassa

import (
	"encoding/json"
	"fmt"
	models "github.com/wDRxxx/yookassa-go-sdk/yookassa/models/refund"
	"net/http"
)

const (
	refundsEndpoint = "/refunds"
)

func (c *Client) CreateRefund(refund *models.Refund) (*models.Refund, error) {
	marshalledRefund, err := json.Marshal(&refund)
	if err != nil {
		return nil, err
	}

	resp, err := c.sendRequest(refundsEndpoint, http.MethodPost, marshalledRefund)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 300 {
		respErr, err := unmarshallResponseError(resp)
		if err != nil {
			return nil, err
		}

		return nil, respErr
	}

	var respRefund models.Refund
	err = unmarshallResponseData(resp, &respRefund)
	if err != nil {
		return nil, err
	}

	return &respRefund, nil
}

func (c *Client) RefundsList(opts *models.RefundsListOptions) (*models.RefundsList, error) {
	marshalledOpts, err := json.Marshal(&opts)
	if err != nil {
		return nil, err
	}

	resp, err := c.sendRequest(refundsEndpoint, http.MethodGet, marshalledOpts)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 300 {
		respErr, err := unmarshallResponseError(resp)
		if err != nil {
			return nil, err
		}

		return nil, respErr
	}

	var respRefunds models.RefundsList
	err = unmarshallResponseData(resp, &respRefunds)
	if err != nil {
		return nil, err
	}

	return &respRefunds, nil
}

func (c *Client) RefundInfo(refundID string) (*models.Refund, error) {
	endpoint := fmt.Sprintf("%s/%s", refundsEndpoint, refundID)

	resp, err := c.sendRequest(endpoint, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode > 300 {
		respErr, err := unmarshallResponseError(resp)
		if err != nil {
			return nil, err
		}

		return nil, respErr
	}

	var respRefund models.Refund
	err = unmarshallResponseData(resp, &respRefund)
	if err != nil {
		return nil, err
	}

	return &respRefund, nil
}
