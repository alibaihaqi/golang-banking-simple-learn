package service

import (
	"github.com/alibaihaqi/banking/domain"
	"github.com/alibaihaqi/banking/dto"
	"github.com/alibaihaqi/banking/errs"
	"time"
)

type TransactionService interface {
	NewTransaction(request *dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError)
}

type DefaultTransactionService struct {
	aRepo domain.AccountRepository
	tRepo domain.TransactionRepository
}

func (s DefaultTransactionService) NewTransaction(req *dto.NewTransactionRequest) (*dto.NewTransactionResponse, *errs.AppError) {
	err := req.ValidateTransactionType()
	if err != nil {
		return nil, err
	}

	tr := domain.Transaction{
		TransactionId:   "",
		AccountId:       req.AccountId,
		Amount:          req.Amount,
		TransactionType: req.TransactionType,
		TransactionDate: time.Now().Format("2006-01-02 15:04:05"),
	}

	// Check if balance is still available to be withdraw
	if tr.TransactionType == "withdrawal" {
		acc, err := s.aRepo.GetBalance(req.AccountId)

		if err != nil {
			return nil, err
		}

		if acc.Amount < tr.Amount {
			return nil, errs.NewValidationError("Balance is not enough to withdraw")
		}
	}

	// Create a transaction
	transaction, err := s.tRepo.NewTransaction(tr)
	if err != nil {
		return nil, errs.NewInternalSystemError(err.Message)
	}

	// Update Balance to your accounts
	updateAccount, appError := s.aRepo.UpdateAccount(*transaction)
	if appError != nil {
		return nil, err
	}
	rDto := dto.NewTransactionResponse{
		TransactionId: tr.TransactionId,
		Amount:        updateAccount.Amount,
	}

	return &rDto, nil
}

func NewTransactionService(a domain.AccountRepository, t domain.TransactionRepository) DefaultTransactionService {
	return DefaultTransactionService{
		aRepo: a,
		tRepo: t,
	}
}
