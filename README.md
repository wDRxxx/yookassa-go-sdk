# Yookassa Go SDK

## Installation
`go get github.com/wDRxxx/yookassa-go-sdk`
## Usage example
Use sdk according to [official documentation](https://yookassa.ru/developers/api?codeLang=bash&lang=ru)
```go
package main

import (
	"github.com/wDRxxx/yookassa-go-sdk/yookassa"
	"github.com/wDRxxx/yookassa-go-sdk/yookassa/models"
	"github.com/wDRxxx/yookassa-go-sdk/yookassa/models/payment"
	"log"
)

func main() {
	c := yookassa.NewClient("shopID", "shopKey")

	payment, err := c.CreatePayment(&payment.Payment{
		Amount: &models.Amount{
			Value:    "1000.00",
			Currency: "RUB",
		},
		PaymentMethodData: &payment.PaymentMethodData{
			Type: "bank_card",
		},
		Confirmation: &models.Confirmation{
			Type:      "redirect",
			ReturnURL: "https://google.com",
		},
		Receipt: &models.Receipt{
			Email: "email@email.com",
			Phone: "88005553535",
			Items: []*models.Item{
				{
					Description: "some description",
					Amount: &models.Amount{
						Value:    "1000.00",
						Currency: "RUB",
					},
					Quantity: 1,
					VATCode:  1,
				},
			},
		},
	})
	if err != nil {
		log.Println(err.Error())
		return
	}

	log.Println(payment)
}
```