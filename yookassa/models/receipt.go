package models

type Recipient struct {
	AccountID string `json:"account_id,omitempty"`
	GatewayID string `json:"gateway_id,omitempty"`
}

type Receipt struct {
	Customer                  *Customer                  `json:"customer,omitempty"`
	Items                     []*Item                    `json:"items,omitempty"`
	Email                     string                     `json:"email,omitempty"`
	Phone                     string                     `json:"phone,omitempty"`
	TaxSystemCode             *int                       `json:"tax_system_code,omitempty"`
	ReceiptIndustryDetails    []*ReceiptIndustryDetail   `json:"receipt_industry_details,omitempty"`
	ReceiptOperationalDetails *ReceiptOperationalDetails `json:"receipt_operational_details,omitempty"`
}

type Customer struct {
	FullName string `json:"full_name,omitempty"`
	INN      string `json:"inn,omitempty"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
}

type Item struct {
	Description                   string                         `json:"description,omitempty"`
	Amount                        *Amount                        `json:"amount,omitempty"`
	VATCode                       int                            `json:"vat_code,omitempty"`
	Quantity                      float64                        `json:"quantity,omitempty"`
	Measure                       string                         `json:"measure,omitempty"`
	MarkQuantity                  *MarkQuantity                  `json:"mark_quantity,omitempty"`
	PaymentSubject                string                         `json:"payment_subject,omitempty"`
	PaymentMode                   string                         `json:"payment_mode,omitempty"`
	CountryOfOriginCode           string                         `json:"country_of_origin_code,omitempty"`
	CustomsDeclarationNumber      string                         `json:"customs_declaration_number,omitempty"`
	Excise                        string                         `json:"excise,omitempty"`
	ProductCode                   string                         `json:"product_code,omitempty"`
	MarkCodeInfo                  *MarkCodeInfo                  `json:"mark_code_info,omitempty"`
	MarkMode                      string                         `json:"mark_mode,omitempty"`
	PaymentSubjectIndustryDetails *PaymentSubjectIndustryDetails `json:"payment_subject_industry_details,omitempty"`
}

type PaymentSubjectIndustryDetails struct {
	FederalID      string `json:"federal_id,omitempty"`
	DocumentDate   string `json:"document_date,omitempty"`
	DocumentNumber string `json:"document_number,omitempty"`
	Value          string `json:"value,omitempty"`
}

type MarkQuantity struct {
	Numerator   int `json:"numerator,omitempty"`
	Denominator int `json:"denominator,omitempty"`
}

type MarkCodeInfo struct {
	MarkCodeRaw string `json:"mark_code_raw,omitempty"`
	Unknown     string `json:"unknown,omitempty"`
	EAN8        string `json:"ean_8,omitempty"`
	EAN13       string `json:"ean_13,omitempty"`
	ITF14       string `json:"itf_14,omitempty"`
	GS10        string `json:"gs_10,omitempty"`
	GS1M        string `json:"gs_1m,omitempty"`
	Short       string `json:"short,omitempty"`
	Fur         string `json:"fur,omitempty"`
	Egais20     string `json:"egais_20,omitempty"`
	Egais30     string `json:"egais_30,omitempty"`
}

type ReceiptIndustryDetail struct {
	FederalID      string `json:"federal_id,omitempty"`
	DocumentDate   string `json:"document_date,omitempty"`
	DocumentNumber string `json:"document_number,omitempty"`
	Value          string `json:"value,omitempty"`
}

type ReceiptOperationalDetails struct {
	OperationID int    `json:"operation_id,omitempty"`
	Value       string `json:"value,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
}
