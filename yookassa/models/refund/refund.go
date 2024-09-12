package refund

import (
	models "github.com/wDRxxx/yookassa-go-sdk/yookassa/models"
)

type Refund struct {
	ID                  string                      `json:"id,omitempty"`
	PaymentID           string                      `json:"payment_id"`
	Status              string                      `json:"status"`
	CancellationDetails *models.CancellationDetails `json:"cancellation_details,omitempty"`
	ReceiptRegistration string                      `json:"receipt_registration,omitempty"`
	CreatedAt           string                      `json:"created_at"`
	Amount              *models.Amount              `json:"amount"`
	Description         string                      `json:"description,omitempty"`
	Sources             []*models.Source            `json:"sources,omitempty"`
	Deal                *models.Deal                `json:"deal,omitempty"`
	RefundMethod        *RefundMethod               `json:"refund_method,omitempty"`
	Receipt             *models.Receipt             `json:"receipt,omitempty"`
	RefundMethodData    *RefundMethodData           `json:"refund_method_data,omitempty"`
}

type RefundMethod struct {
	Type                  string                        `json:"type,omitempty"`
	SbpOperationID        string                        `json:"sbp_operation_id,omitempty"`
	Articles              []*models.Article             `json:"articles,omitempty"`
	ElectronicCertificate *models.ElectronicCertificate `json:"electronic_certificate,omitempty"`
}

type RefundMethodData struct {
	Type                  string                        `json:"type,omitempty"`
	Articles              []*models.Article             `json:"articles,omitempty"`
	ElectronicCertificate *models.ElectronicCertificate `json:"electronic_certificate,omitempty"`
}

type RefundsListOptions struct {
	CreatedAtGte  string `json:"created_at.gte,omitempty"`
	CreatedAtGt   string `json:"created_at.gt,omitempty"`
	CreatedAtLte  string `json:"created_at.lte,omitempty"`
	CreatedAtLt   string `json:"created_at.lt,omitempty"`
	CapturedAtGte string `json:"captured_at.gte,omitempty"`
	CapturedAtGt  string `json:"captured_at.gt,omitempty"`
	CapturedAtLte string `json:"captured_at.lte,omitempty"`
	CapturedAtLt  string `json:"captured_at.lt,omitempty"`
	PaymentID     string `json:"payment_id,omitempty"`
	Status        string `json:"status,omitempty"`
	Limit         int    `json:"limit,omitempty"`
	Cursor        string `json:"cursor,omitempty"`
}

type RefundsList struct {
	Type       string  `json:"type,omitempty"`
	Items      *Refund `json:"items,omitempty"`
	NextCursor string  `json:"next_cursor,omitempty"`
}
