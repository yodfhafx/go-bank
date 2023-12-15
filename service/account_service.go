package service

import (
	"strings"
	"time"

	"github.com/yodfhafx/go-bank/errs"
	"github.com/yodfhafx/go-bank/logs"
	"github.com/yodfhafx/go-bank/repository"
)

type accountService struct {
	accRepo repository.AccountRepository
}

func NewAccountService(accRepo repository.AccountRepository) AccountService {
	return accountService{accRepo: accRepo}
}

func (s accountService) NewAccount(customerID int, request NewAccountRequest) (*AccountResponse, error) {
	//Validate
	if request.Amount < 1000 {
		return nil, errs.NewValidationError("Amount must at least 1,000")
	}
	if strings.ToLower(request.AccountType) != "saving" && strings.ToLower(request.AccountType) != "checking" {
		return nil, errs.NewValidationError("Account type should be saving or checking")
	}

	account := repository.Account{
		CustomerID:  customerID,
		OpeningDate: time.Now().Format("2006-1-2 15:04:05"),
		AccountType: request.AccountType,
		Amount:      request.Amount,
		Status:      1,
	}

	newAcc, err := s.accRepo.Create(account)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	accResponse := AccountResponse{
		AccountID:   newAcc.AccountID,
		OpeningDate: newAcc.OpeningDate,
		AccountType: newAcc.AccountType,
		Amount:      newAcc.Amount,
		Status:      newAcc.Status,
	}

	return &accResponse, nil
}

func (s accountService) GetAccounts(customerID int) ([]AccountResponse, error) {
	accounts, err := s.accRepo.GetAll(customerID)
	if err != nil {
		logs.Error(err)
		return nil, errs.NewUnexpectedError()
	}

	accResponses := []AccountResponse{}
	for _, account := range accounts {
		accResponses = append(accResponses, AccountResponse{
			AccountID:   account.AccountID,
			OpeningDate: account.OpeningDate,
			AccountType: account.AccountType,
			Amount:      account.Amount,
			Status:      account.Status,
		})
	}
	return accResponses, nil
}
