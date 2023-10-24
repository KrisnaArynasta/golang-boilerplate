package Model

import "time"

type EnumDepositResponseStatus string

type DepositTransaction struct {
	OrderId                   string                    `json:"orderId"`
	TransactionId             int                       `json:"transactionId"`
	DateCreated               time.Time                 `json:"dateCreated"`
	DepositMethod             string                    `json:"depositMethod"`
	Processor                 string                    `json:"processor"`
	Amount                    float64                   `json:"amount"`
	Currency                  string                    `json:"currency"`
	Status                    EnumDepositResponseStatus `json:"status"`
	DateUpdated               time.Time                 `json:"dateUpdated"`
	DepositorUserId           string                    `json:"depositorUserId"`
	Notes                     string                    `json:"notes"`
	Channel                   string                    `json:"channel"`
	MessageAuthenticationCode string                    `json:"messageAuthenticationCode"`
	ProcessingCurrency        string                    `json:"processingCurrency"`
	ProcessingAmount          float64                   `json:"processingAmount"`
}

type PaymentPageSession struct {
	SessionToken   string    `json:"sessionToken"`
	Expires        time.Time `json:"expires"`
	SessionType    string    `json:"sessionType"`
	OrderId        string    `json:"orderId"`
	PaymentPageUrl string    `json:"paymentPageUrl"`
}

type RequestDepositResponse struct {
	DepositTransaction        DepositTransaction `json:"depositTransaction"`
	PaymentPageSession        PaymentPageSession `json:"paymentPageSession"`
	Success                   bool               `json:"success"`
	OrderId                   string             `json:"orderId"`
	MessageAuthenticationCode string             `json:"messageAuthenticationCode"`
	Correlation               string             `json:"correlation"`
	Message                   string             `json:"message"`
}
