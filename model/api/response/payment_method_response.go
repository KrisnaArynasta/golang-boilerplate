package Model

type PaymentMethodResponse []PaymentMethodElement

type PaymentMethodElement struct {
	Method  string   `json:"method"`
	Details []Detail `json:"details"`
}

type Detail struct {
	Currency string  `json:"currency"`
	Limits   []Limit `json:"limits"`
}

type Limit struct {
	MinTier              int64   `json:"minTier"`
	MaxTier              int64   `json:"maxTier"`
	MinTransactionAmount float64 `json:"minTransactionAmount"`
	MaxTransactionAmount float64 `json:"maxTransactionAmount"`
	BankBased            bool    `json:"bankBased"`
	Banks                []Bank  `json:"banks,omitempty"`
}

type Bank struct {
	BankCode    string `json:"bankCode"`
	BankName    string `json:"bankName"`
	EnglishName string `json:"englishName"`
}
