package payment

import (
	"github.com/wDRxxx/yookassa-go-sdk/yookassa/models"
)

type Passenger struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

type Leg struct {
	DepartureAirport   string `json:"departure_airport"`
	DestinationAirport string `json:"destination_airport"`
	DepartureDate      string `json:"departure_date"`
	CarrierCode        string `json:"carrier_code,omitempty"`
}

type Airline struct {
	TicketNumber     string       `json:"ticket_number,omitempty"`
	BookingReference string       `json:"booking_reference,omitempty"`
	Passengers       []*Passenger `json:"passengers,omitempty"`
	Legs             []*Leg       `json:"legs,omitempty"`
}

type PaymentRequest struct {
	Airline   *Airline           `json:"airline,omitempty"`
	Transfers []*models.Transfer `json:"transfers,omitempty"`
	Metadata  map[string]string  `json:"metadata,omitempty"`
}
