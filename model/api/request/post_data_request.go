package Model

type PostDataRequest struct {
	OrderId                   string `json:"orderId"`
	Amount                    string `json:"amount"`
	Currency                  string `json:"currency"`
	DateTime                  string `json:"datetime"`
	Language                  string `json:"language"`
	DepositorUserId           string `json:"depositoruserid"`
	DepositMethod             string `json:"depositmethod"`
	CallbackUrl               string `json:"callbackurl"`
	RedirectUrl               string `json:"redirecturl"`
	CancelUrl                 string `json:"cancelurl"`
	MessageAuthenticationCode string `json:"messageauthenticationcode"`
	DepositorBank             string `json:"depositorbank"`
	DepositorName             string `json:"depositorname"`
}
