package receipt

import (
	"github.com/wDRxxx/yookassa-go-sdk/yookassa/models"
)

type Receipt struct {
	ID                        string                            `json:"id,omitempty"`
	Type                      string                            `json:"type,omitempty"`
	PaymentID                 string                            `json:"payment_id,omitempty"`
	RefundID                  string                            `json:"refund_id,omitempty"`
	Status                    string                            `json:"status,omitempty"`
	FiscalDocumentNumber      string                            `json:"fiscal_document_number,omitempty"`
	FiscalStorageNumber       string                            `json:"fiscal_storage_number,omitempty"`
	FiscalAttribute           string                            `json:"fiscal_attribute,omitempty"`
	RegisteredAt              string                            `json:"registered_at,omitempty"`
	FiscalProviderID          string                            `json:"fiscal_provider_id,omitempty"`
	Items                     []*models.Item                    `json:"items,omitempty"`
	Settlements               []*models.Settlement              `json:"settlements,omitempty"`
	OnBehalfOf                string                            `json:"on_behalf_of,omitempty"`
	TaxSystemCode             int                               `json:"tax_system_code,omitempty"`
	ReceiptIndustryDetails    *[]models.ReceiptIndustryDetail   `json:"receipt_industry_details,omitempty"`
	ReceiptOperationalDetails *models.ReceiptOperationalDetails `json:"receipt_operational_details,omitempty"`
	Customer                  *models.Customer                  `json:"customer,omitempty"`
	Send                      bool                              `json:"send,omitempty"`
	AdditionalUserProps       *models.AdditionalUserProps       `json:"additional_user_props,omitempty"`
}

type ReceiptsListOptions struct {
	CreatedAtGte  string `json:"created_at.gte,omitempty"`
	CreatedAtGt   string `json:"created_at.gt,omitempty"`
	CreatedAtLte  string `json:"created_at.lte,omitempty"`
	CreatedAtLt   string `json:"created_at.lt,omitempty"`
	CapturedAtGte string `json:"captured_at.gte,omitempty"`
	CapturedAtGt  string `json:"captured_at.gt,omitempty"`
	CapturedAtLte string `json:"captured_at.lte,omitempty"`
	CapturedAtLt  string `json:"captured_at.lt,omitempty"`
	PaymentID     string `json:"payment_id,omitempty"`
	RefundID      string `json:"refund_id,omitempty"`
	Limit         int    `json:"limit,omitempty"`
	Cursor        string `json:"cursor,omitempty"`
}
