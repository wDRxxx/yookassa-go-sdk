package payment

import (
	"github.com/wDRxxx/yookassa-go-sdk/yookassa/models"
)

type PayerBankDetails struct {
	BankID     string `json:"bank_id,omitempty"`
	BIC        string `json:"bic,omitempty"`
	FullName   string `json:"full_name,omitempty"`
	ShortName  string `json:"short_name,omitempty"`
	INN        string `json:"inn,omitempty"`
	BankName   string `json:"bank_name,omitempty"`
	BankBranch string `json:"bank_branch,omitempty"`
	BankBIK    string `json:"bank_bik,omitempty"`
	Account    string `json:"account,omitempty"`
	KPP        string `json:"kpp,omitempty"`
}

type CardProduct struct {
	Code string `json:"code"`
	Name string `json:"name,omitempty"`
}

type Card struct {
	First6        string       `json:"first_6,omitempty"`
	Last4         string       `json:"last_4"`
	ExpiryYear    string       `json:"expiry_year"`
	ExpiryMonth   string       `json:"expiry_month"`
	CardType      string       `json:"card_type"`
	CardProduct   *CardProduct `json:"card_product,omitempty"`
	IssuerCountry string       `json:"issuer_country,omitempty"`
	IssuerName    string       `json:"issuer_name,omitempty"`
	Source        string       `json:"source,omitempty"`
}

type PaymentMethod struct {
	ID               string            `json:"id,omitempty"`
	Type             string            `json:"type"`
	Saved            bool              `json:"saved,omitempty"`
	Title            string            `json:"title,omitempty"`
	PayerBankDetails *PayerBankDetails `json:"payer_bank_details,omitempty"`
	DiscountAmount   *models.Amount    `json:"discount_amount,omitempty"`
	LoanOption       string            `json:"loan_option,omitempty"`
	Login            string            `json:"login,omitempty"`
	Articles         []*models.Article `json:"articles,omitempty"`
	AccountNumber    string            `json:"account_number,omitempty"`
	Phone            string            `json:"phone,omitempty"`
}
