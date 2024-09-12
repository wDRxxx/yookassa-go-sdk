package yookassa

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ResponseError struct {
	Type        string `json:"type"`
	ID          string `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Parameter   string `json:"parameter,omitempty"`
}

func (e *ResponseError) Error() string {
	if e.Parameter != "" {
		return fmt.Sprintf("%s: %s - %s", e.Code, e.Description, e.Parameter)
	}
	return fmt.Sprintf("%s: %s", e.Code, e.Description)
}

func unmarshallResponseData(resp *http.Response, data any) error {
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, &data)
	if err != nil {
		return err
	}

	return nil
}

func unmarshallResponseError(resp *http.Response) (*ResponseError, error) {
	var respError ResponseError
	err := unmarshallResponseData(resp, &respError)
	if err != nil {
		return nil, err
	}

	return &respError, nil
}
