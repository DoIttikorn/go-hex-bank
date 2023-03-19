package service

type AccountRequest struct {
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

type AccountResponse struct {
	AccountID   int     `json:"account_id"`
	OpeningDate string  `json:"opening_date"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
	Status      string  `json:"status"`
}

type AccountService interface {
	ExecuteNewAccount(account AccountRequest) (*AccountResponse, error)
	ExecuteGetAccount(int) ([]AccountResponse, error)
}
