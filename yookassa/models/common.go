package models

type Amount struct {
	Value    string `json:"value"`
	Currency string `json:"currency"`
}

type Confirmation struct {
	Type              string `json:"type"`
	ConfirmationToken string `json:"confirmation_token,omitempty"`
	ConfirmationURL   string `json:"confirmation_url,omitempty"`
	ConfirmationData  string `json:"confirmation_data,omitempty"`
	Enforce           bool   `json:"enforce,omitempty"`
	ReturnURL         string `json:"return_url,omitempty"`
}

type CancellationDetails struct {
	Party  string `json:"party"`
	Reason string `json:"reason"`
}

type ThreeDSecure struct {
	Applied bool `json:"applied"`
}

type AuthorizationDetails struct {
	RRN          string        `json:"rrn,omitempty"`
	AuthCode     string        `json:"auth_code,omitempty"`
	ThreeDSecure *ThreeDSecure `json:"three_d_secure"`
}

type Transfer struct {
	AccountID         string            `json:"account_id"`
	Amount            *Amount           `json:"amount"`
	Status            string            `json:"status"`
	PlatformFeeAmount *Amount           `json:"platform_fee_amount,omitempty"`
	Description       string            `json:"description"`
	Metadata          map[string]string `json:"metadata"`
}

type Settlement struct {
	Type   string  `json:"type"`
	Amount *Amount `json:"amount"`
}

type Deal struct {
	ID                string        `json:"id"`
	Settlements       []*Settlement `json:"settlements,omitempty"`
	RefundSettlements []*Settlement `json:"refund_settlements,omitempty"`
}

type Receiver struct {
	Type          string `json:"type"`
	AccountNumber string `json:"account_number,omitempty"`
	BIC           string `json:"bic,omitempty,omitempty"`
}

type Source struct {
	AccountID         string  `json:"account_id"`
	Amount            *Amount `json:"amount"`
	PlatformFeeAmount *Amount `json:"platform_fee_amount"`
}
type Article struct {
	ArticleNumber         int            `json:"article_number,omitempty"`
	TruCode               string         `json:"tru_code,omitempty"`
	ArticleCode           string         `json:"article_code,omitempty"`
	ArticleName           string         `json:"article_name,omitempty"`
	Certificates          []*Certificate `json:"certificates,omitempty"`
	TruQuantity           int            `json:"tru_quantity,omitempty"`
	AvailableCompensation *Amount        `json:"available_compensation,omitempty"`
	AppliedCompensation   *Amount        `json:"applied_compensation,omitempty"`
	Quantity              int            `json:"quantity,omitempty"`
	Price                 *Amount        `json:"price,omitempty"`
}

type Certificate struct {
	CertificateID string `json:"certificate_id"`
	TruQuantity   int    `json:"tru_quantity"`
}

type ElectronicCertificate struct {
	Amount   *Amount `json:"amount"`
	BasketID string  `json:"basket_id"`
}

type AdditionalUserProps struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
