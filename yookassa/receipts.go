package yookassa

import (
	"encoding/json"
	"fmt"
	models "github.com/wDRxxx/yookassa-go-sdk/yookassa/models/receipt"
	"net/http"
)

const (
	receiptsEndpoint = "/receipts"
)

func (c *Client) CreateReceipt(refund *models.Receipt) (*models.Receipt, error) {
	marshalledReceipt, err := json.Marshal(&refund)
	if err != nil {
		return nil, err
	}

	resp, err := c.sendRequest(receiptsEndpoint, http.MethodPost, marshalledReceipt)
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

	var respReceipt models.Receipt
	err = unmarshallResponseData(resp, &respReceipt)
	if err != nil {
		return nil, err
	}

	return &respReceipt, nil
}

func (c *Client) ReceiptsList(opts *models.ReceiptsListOptions) (*models.ReceiptsListOptions, error) {
	marshalledOpts, err := json.Marshal(&opts)
	if err != nil {
		return nil, err
	}

	resp, err := c.sendRequest(receiptsEndpoint, http.MethodGet, marshalledOpts)
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

	var respReceipts models.ReceiptsListOptions
	err = unmarshallResponseData(resp, &respReceipts)
	if err != nil {
		return nil, err
	}

	return &respReceipts, nil
}

func (c *Client) ReceiptInfo(refundID string) (*models.Receipt, error) {
	endpoint := fmt.Sprintf("%s/%s", receiptsEndpoint, refundID)

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

	var respReceipt models.Receipt
	err = unmarshallResponseData(resp, &respReceipt)
	if err != nil {
		return nil, err
	}

	return &respReceipt, nil
}
