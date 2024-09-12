package yookassa

import (
	"encoding/json"
	"fmt"
	models "github.com/wDRxxx/yookassa-go-sdk/yookassa/models/payment"
	"net/http"
)

const (
	paymentsEndpoint = "/payments"
)

func (c *Client) CreatePayment(payment *models.Payment) (*models.Payment, error) {
	marshalledPayment, err := json.Marshal(&payment)
	if err != nil {
		return nil, err
	}

	resp, err := c.sendRequest(paymentsEndpoint, http.MethodPost, marshalledPayment)
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

	var respPayment models.Payment
	err = unmarshallResponseData(resp, &respPayment)
	if err != nil {
		return nil, err
	}

	return &respPayment, nil
}

func (c *Client) PaymentsList(opts *models.PaymentsListOptions) (*models.PaymentsList, error) {
	marshalledOpts, err := json.Marshal(&opts)
	if err != nil {
		return nil, err
	}

	resp, err := c.sendRequest(paymentsEndpoint, http.MethodGet, marshalledOpts)
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

	var payments models.PaymentsList
	err = unmarshallResponseData(resp, &payments)

	if err != nil {
		return nil, err
	}

	return &payments, nil
}

func (c *Client) PaymentInfo(paymentID string) (*models.Payment, error) {
	endpoint := fmt.Sprintf("%s/%s", paymentsEndpoint, paymentID)

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

	var payment models.Payment
	err = unmarshallResponseData(resp, &payment)

	if err != nil {
		return nil, err
	}

	return &payment, nil
}

func (c *Client) CapturePayment(payment *models.Payment) (*models.Payment, error) {
	marshalledPayment, err := json.Marshal(&payment)
	if err != nil {
		return nil, err
	}

	endpoint := fmt.Sprintf("%s/%s/capture", paymentsEndpoint, payment.ID)
	resp, err := c.sendRequest(endpoint, http.MethodPost, marshalledPayment)
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

	var respPayment models.Payment
	err = unmarshallResponseData(resp, &respPayment)
	if err != nil {
		return nil, err
	}

	return &respPayment, nil
}

func (c *Client) CancelPayment(paymentID string) (*models.Payment, error) {
	endpoint := fmt.Sprintf("%s/%s/cancel", paymentsEndpoint, paymentID)
	resp, err := c.sendRequest(endpoint, "POST", nil)
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

	var respPayment models.Payment
	err = unmarshallResponseData(resp, &respPayment)
	if err != nil {
		return nil, err
	}

	return &respPayment, nil
}
