package domain

import (
	"github.com/alibaihaqi/banking/dto"
	"github.com/alibaihaqi/banking/errs"
)

type Account struct {
	AccountId   string  `json:"account_id" db:"account_id"`
	CustomerId  string  `json:"customer_id" db:"customer_id"`
	OpeningDate string  `json:"opening_date" db:"opening_date"`
	AccountType string  `json:"account_type" db:"account_type"`
	Amount      float64 `json:"amount" db:"amount"`
	Status      string  `json:"status" db:"status"`
}

type AccountRepository interface {
	GetBalance(string) (*Account, *errs.AppError)
	Save(Account) (*Account, *errs.AppError)
	UpdateAccount(Transaction) (*Account, *errs.AppError)
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{
		AccountId:   a.AccountId,
		CustomerId:  a.CustomerId,
		AccountType: a.AccountType,
		Amount:      a.Amount,
	}
}
