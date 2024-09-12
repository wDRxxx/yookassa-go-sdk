package payment

import (
	"github.com/wDRxxx/yookassa-go-sdk/yookassa/models"
)

type Payment struct {
	ID                   string                       `json:"id,omitempty"`
	Status               string                       `json:"status,omitempty"`
	Amount               *models.Amount               `json:"amount,omitempty"`
	IncomeAmount         *models.Amount               `json:"income_amount,omitempty"`
	Description          string                       `json:"description,omitempty"`
	Recipient            *models.Recipient            `json:"recipient,omitempty"`
	PaymentMethod        *PaymentMethod               `json:"payment_method,omitempty"`
	CapturedAt           string                       `json:"captured_at,omitempty"`
	ExpiresAt            string                       `json:"expires_at,omitempty"`
	Test                 bool                         `json:"test,omitempty"`
	RefundAmount         *models.Amount               `json:"refund_amount,omitempty"`
	Paid                 bool                         `json:"paid,omitempty"`
	Refundable           bool                         `json:"refundable,omitempty"`
	ReceiptRegistration  string                       `json:"receipt_registration,omitempty"`
	Metadata             map[string]string            `json:"metadata,omitempty,omitempty"`
	CancellationDetails  *models.CancellationDetails  `json:"cancellation_details,omitempty"`
	AuthorizationDetails *models.AuthorizationDetails `json:"authorization_details,omitempty"`
	Transfers            []*models.Transfer           `json:"transfers,omitempty"`
	Deal                 *models.Deal                 `json:"deal,omitempty"`
	MerchantCustomerID   string                       `json:"merchant_customer_id,omitempty"`
	Receipt              *models.Receipt              `json:"receipt,omitempty"`
	PaymentToken         string                       `json:"payment_token,omitempty"`
	PaymentMethodData    *PaymentMethodData           `json:"payment_method_data,omitempty"`
	PaymentTokenID       string                       `json:"payment_token_id,omitempty"`
	Confirmation         *models.Confirmation         `json:"confirmation,omitempty"`
	SavePaymentMethod    bool                         `json:"save_payment_method,omitempty"`
	Capture              bool                         `json:"capture,omitempty"`
	ClientIP             string                       `json:"client_ip,omitempty"`
	Airline              *Airline                     `json:"airline,omitempty"`
	Receiver             *models.Receiver             `json:"receiver,omitempty"`
}

type PaymentMethodData struct {
	Type                  string                        `json:"type"`
	Card                  *Card                         `json:"card,omitempty"`
	Phone                 string                        `json:"phone,omitempty"`
	PaymentPurpose        string                        `json:"payment_purpose,omitempty"`
	VatData               *VatData                      `json:"vat_data,omitempty"`
	Articles              []*models.Article             `json:"articles,omitempty"`
	ElectronicCertificate *models.ElectronicCertificate `json:"electronic_certificate"`
}

type VatData struct {
	Type   string         `json:"type"`
	Amount *models.Amount `json:"amount,omitempty"`
	Rate   string         `json:"rate,omitempty"`
}

type PaymentsList struct {
	Type       string   `json:"type,omitempty"`
	Items      *Payment `json:"items,omitempty"`
	NextCursor string   `json:"next_cursor,omitempty"`
}

type PaymentsListOptions struct {
	CreatedAtGte  string `json:"created_at.gte,omitempty"`
	CreatedAtGt   string `json:"created_at.gt,omitempty"`
	CreatedAtLte  string `json:"created_at.lte,omitempty"`
	CreatedAtLt   string `json:"created_at.lt,omitempty"`
	CapturedAtGte string `json:"captured_at.gte,omitempty"`
	CapturedAtGt  string `json:"captured_at.gt,omitempty"`
	CapturedAtLte string `json:"captured_at.lte,omitempty"`
	CapturedAtLt  string `json:"captured_at.lt,omitempty"`
	PaymentMethod string `json:"payment_method,omitempty"`
	Status        string `json:"status,omitempty"`
	Limit         int    `json:"limit,omitempty"`
	Cursor        string `json:"cursor,omitempty"`
}
