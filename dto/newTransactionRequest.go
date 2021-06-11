package dto

import (
	"github.com/alibaihaqi/banking/errs"
	"strings"
)

type NewTransactionRequest struct {
	AccountId       string  `json:"account_id"`
	TransactionType string  `json:"transaction_type"`
	Amount          float64 `json:"amount"`
}

func (r NewTransactionRequest) ValidateTransactionType() *errs.AppError {
	if strings.ToLower(r.TransactionType) != "withdrawal" && strings.ToLower(r.TransactionType) != "deposit" {
		return errs.NewValidationError("Transaction type should be withdrawal or deposit")
	}
	if r.Amount < 0 {
		return errs.NewValidationError("Amount must be positive value")
	}
	if r.AccountId == "" {
		return errs.NewValidationError("Account Id cannot be empty")
	}
	return nil
}
