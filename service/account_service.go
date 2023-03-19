package service

import "github.com/Doittikorn/bank/repository"

type accountService struct {
	accountRepo repository.AccountRepository
}

func NewAccountService(accountRepo repository.AccountRepository) AccountService {
	return &accountService{accountRepo: accountRepo}
}

func (s *accountService) ExecuteNewAccount(account AccountRequest) (*AccountResponse, error) {
	return nil, nil
}

func (s *accountService) ExecuteGetAccount(customerID int) ([]AccountResponse, error) {
	account, err := s.accountRepo.GetAll(customerID)

	if err != nil {
		return nil, err
	}

	var accounts []AccountResponse
	for _, a := range account {
		accounts = append(accounts, AccountResponse{
			AccountID:   a.AccountID,
			OpeningDate: a.OpeningDate,
			AccountType: a.AccountType,
			Amount:      a.Amount,
			Status:      a.Status,
		})
	}

	return accounts, nil
}
