package domain

import (
	"github.com/alibaihaqi/banking/dto"
	"github.com/alibaihaqi/banking/errs"
)

type Transaction struct {
	TransactionId   string
	AccountId       string
	Amount          float64
	TransactionType string
	TransactionDate string
}

type TransactionRepository interface {
	NewTransaction(Transaction) (*Transaction, *errs.AppError)
}

func (t Transaction) ToNewTransactionResponseDto() dto.NewTransactionResponse {
	return dto.NewTransactionResponse{
		TransactionId: t.TransactionId,
		Amount:        t.Amount,
	}
}
